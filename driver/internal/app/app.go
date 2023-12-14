package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/notifier"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/repository"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/server"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/runner"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"os"
	"path"
)

const (
	swaggerFileName = "swagger.json"
)

type App struct {
	runner.RunStopper
	cfg *Config

	driverServer    *server.DriverServer
	httpProxyServer *http.Server

	isStopping bool
}

func NewApp(config *Config, runStopperPreset runner.RunStopper) *App {
	return &App{
		RunStopper: runStopperPreset,
		cfg:        config,
		isStopping: false,
	}
}

func (a *App) Run(ctx context.Context) error {
	err := a.RunStopper.Run(ctx)
	if err != nil {
		return err
	}

	db, err := repository.NewMongoDB(ctx, repository.MongoConfig{
		URI:        a.cfg.Mongo.URI,
		AuthSource: a.cfg.Mongo.AuthSource,
		Username:   os.Getenv("MONGO_USER"),
		Password:   os.Getenv("MONGO_PASSWORD"),
	})
	if err != nil {
		return fmt.Errorf("can't init mongo db: %w", err)
	}
	defer db.Disconnect(ctx)

	driverRepo := repository.NewDriverRepository(db)
	driverSvc := service.NewDriverService(driverRepo)

	errCh := make(chan error, 2)
	go func() {
		err = a.runGRPCServer(driverSvc)
		errCh <- fmt.Errorf("can't run grpc server: %w", err)
	}()

	go func() {
		err = a.runHTTPServer(ctx, driverSvc)
		errCh <- fmt.Errorf("can't run http proxy server: %w", err)
	}()

	for i := 0; i < cap(errCh); i++ {
		err = <-errCh
		if err != nil && !a.isStopping {
			return err
		}
	}

	return nil
}

func (a *App) runGRPCServer(service *service.DriverService) error {
	driverServer := server.NewDriverServer(service)

	a.driverServer = driverServer
	a.driverServer.Register()

	grpcAddr := a.cfg.GRPC.Addr
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("can't listen: %w", err)
	}
	defer listener.Close()

	loggy.Infoln("starting grpc server on", grpcAddr)

	if err = driverServer.Serve(listener); err != nil {
		return fmt.Errorf("can't serve: %w", err)
	}

	return nil
}

func (a *App) runHTTPServer(ctx context.Context, service *service.DriverService) error {
	httpAddr := a.cfg.HTTP.Addr
	gwMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}))

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := driver.RegisterDriverHandlerFromEndpoint(ctx, gwMux, a.cfg.GRPC.Addr, opts)
	if err != nil {
		return fmt.Errorf("can't register handler for grpc endpoint: %w", err)
	}

	wsNotifier := notifier.NewWSNotifier(service)

	httpMux := http.NewServeMux()
	httpMux.Handle("/", gwMux)
	httpMux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(a.cfg.Swagger.Path, swaggerFileName))
	})
	httpMux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir(a.cfg.Swagger.Path))))

	httpMux.HandleFunc("/trips", wsNotifier.TripsHandler)

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: httpMux,
	}
	a.httpProxyServer = srv

	loggy.Infoln("starting http server on", httpAddr)
	if err = a.httpProxyServer.ListenAndServe(); err != nil {
		return fmt.Errorf("can't serve: %w", err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.isStopping = true

	return multierr.Combine(
		a.httpProxyServer.Shutdown(ctx),
		a.driverServer.GracefulStop(ctx),
		a.RunStopper.Stop(ctx),
	)
}
