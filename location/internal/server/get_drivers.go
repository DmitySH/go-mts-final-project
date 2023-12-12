package server

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/pkg/api/location"
)

func (l *LocationServer) GetDrivers(context.Context, *location.GetDriversRequest) (*location.GetDriversResponse, error) {
	return nil, nil
}
