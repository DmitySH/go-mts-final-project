package server

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/converters"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/pkg/api/location"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (l *LocationServer) UpdateDriverLocation(ctx context.Context, req *location.UpdateDriverLocationRequest) (*location.UpdateDriverLocationResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	driverID, err := uuid.Parse(req.GetDriverId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid driver id")
	}

	latLng, err := converters.LatLngLiteralFromProto(req.GetLocation())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = l.locationService.UpdateDriverLocation(ctx, driverID, latLng)
	if err != nil {
		loggy.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &location.UpdateDriverLocationResponse{}, nil
}
