package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/server"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/runner"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
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
}

func NewApp(config *Config, runStopperPreset runner.RunStopper) *App {
	return &App{
		RunStopper: runStopperPreset,
		cfg:        config,
	}
}

func (a *App) Run(ctx context.Context) error {
	err := a.RunStopper.Run(ctx)
	if err != nil {
		return err
	}

	errCh := make(chan error, 2)
	go func() {
		err = a.runGRPCServer()
		errCh <- fmt.Errorf("can't run grpc server: %w", err)
	}()

	go func() {
		err = a.runHTTPProxy()
		errCh <- fmt.Errorf("can't run http proxy server: %w", err)
	}()

	for i := 0; i < len(errCh); i++ {
		err = <-errCh
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runGRPCServer() error {
	driverServer := server.NewDriverServer()

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

func (a *App) runHTTPProxy() error {
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

	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := driver.RegisterDriverHandlerFromEndpoint(ctx, gwMux, a.cfg.GRPC.Addr, opts)
	if err != nil {
		return fmt.Errorf("can't register handler for grpc endpoint: %w", err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", gwMux)
	httpMux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(a.cfg.Swagger.Path, swaggerFileName))
	})
	httpMux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir(a.cfg.Swagger.Path))))

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: httpMux,
	}
	a.httpProxyServer = srv

	loggy.Infoln("starting http proxy on", httpAddr)
	if err = a.httpProxyServer.ListenAndServe(); err != nil {
		return fmt.Errorf("can't serve: %w", err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	return multierr.Combine(
		a.httpProxyServer.Shutdown(ctx),
		a.driverServer.GracefulStop(ctx),
		a.RunStopper.Stop(ctx),
	)
}
