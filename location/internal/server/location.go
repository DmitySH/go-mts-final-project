package server

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/pkg/api/location"
	"google.golang.org/grpc"
)

type LocationServer struct {
	location.UnimplementedLocationServer

	*grpc.Server
}

func NewLocationServer() *LocationServer {
	return &LocationServer{}
}

func (l *LocationServer) Register() {
	var opts []grpc.ServerOption
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
