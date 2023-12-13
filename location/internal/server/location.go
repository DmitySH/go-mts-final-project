package server

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/pkg/api/location"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type LocationServer struct {
	location.UnimplementedLocationServer
	*grpc.Server

	locationService *service.LocationService
}

func NewLocationServer(locationService *service.LocationService) *LocationServer {
	return &LocationServer{
		locationService: locationService,
	}
}

func (l *LocationServer) Register() {
	opts := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	l.Server = grpc.NewServer(opts...)
	location.RegisterLocationServer(l.Server, l)
}

func (l *LocationServer) GracefulStop(ctx context.Context) error {
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
