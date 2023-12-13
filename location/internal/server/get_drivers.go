package server

import (
	"context"
	"errors"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/converters"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/pkg/api/location"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LocationServer) GetDrivers(ctx context.Context, req *location.GetDriversRequest) (*location.GetDriversResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	latLng, err := converters.LatLngLiteralFromFloats(req.GetLat(), req.GetLng())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.GetRadius() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid radius")
	}

	drivers, err := l.locationService.GetDrivers(ctx, latLng, req.GetRadius())
	if errors.Is(err, service.ErrEntityNotFound) {
		return nil, status.Error(codes.NotFound, "no drivers here")
	}
	if err != nil {
		loggy.Error(err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &location.GetDriversResponse{Drivers: converters.DriversToProto(drivers)}, nil
}
