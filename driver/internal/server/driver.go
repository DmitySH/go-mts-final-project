package server

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
	"google.golang.org/grpc"
)

type DriverServer struct {
	driver.UnimplementedDriverServer

	*grpc.Server
}

func NewDriverServer() *DriverServer {
	return &DriverServer{}
}

func (l *DriverServer) Register() {
	var opts []grpc.ServerOption
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
