package server

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type DriverServer struct {
	driver.UnimplementedDriverServer
	*grpc.Server

	driverService *service.DriverService
}

func NewDriverServer(driverService *service.DriverService) *DriverServer {
	return &DriverServer{
		driverService: driverService,
	}
}

func (l *DriverServer) Register() {
	opts := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	l.Server = grpc.NewServer(opts...)
	driver.RegisterDriverServer(l.Server, l)
}

func (l *DriverServer) GracefulStop(ctx context.Context) error {
	done := make(chan struct{}, 1)

	go func() {
		l.Server.GracefulStop()
		done <- struct{}{}
		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}
