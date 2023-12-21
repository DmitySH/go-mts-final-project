package server

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
)

const userID = "user_id"

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

func (d *DriverServer) Register() {
	opts := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	d.Server = grpc.NewServer(opts...)
	driver.RegisterDriverServer(d.Server, d)
}

func (d *DriverServer) GracefulStop(ctx context.Context) error {
	done := make(chan struct{}, 1)

	go func() {
		d.Server.GracefulStop()
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
