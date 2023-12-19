package clients

import (
	"context"
	"fmt"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/converters"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/location"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LocationClient struct {
	clientConn *grpc.ClientConn
	grpcClient location.LocationClient
}

func NewLocationClient(addr string) (*LocationClient, error) {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("can't dial: %w", err)
	}

	c := &LocationClient{
		clientConn: conn,
		grpcClient: location.NewLocationClient(conn),
	}

	return c, nil
}

func (l *LocationClient) GetDrivers(ctx context.Context, latLng entity.LatLng, radius float64) ([]entity.Driver, error) {
	resp, err := l.grpcClient.GetDrivers(ctx, &location.GetDriversRequest{
		Lat:    latLng.Lat,
		Lng:    latLng.Lng,
		Radius: radius,
	})
	if err != nil {
		return nil, err
	}

	res := make([]entity.Driver, 0, len(resp.GetDrivers()))

	for _, driver := range resp.GetDrivers() {
		res = append(res, converters.DriverFromProto(driver))
	}

	return res, nil
}

func (l *LocationClient) Disconnect() error {
	return l.clientConn.Close()
}
