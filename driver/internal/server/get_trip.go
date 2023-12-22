package server

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/converters"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (d *DriverServer) GetTrips(ctx context.Context, _ *driver.GetTripsRequest) (*driver.GetTripsResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	trips, err := d.driverService.GetTrips(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &driver.GetTripsResponse{}
	for _, trip := range trips {
		resp.Trips = append(resp.Trips, converters.TripFromEntity(trip))
	}

	return resp, nil
}
func (d *DriverServer) GetTripByID(ctx context.Context, req *driver.GetTripByIDRequest) (*driver.GetTripByIDResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	if _, err := uuid.Parse(req.GetTripId()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid trip id format")
	}

	trip, err := d.driverService.GetTripByID(ctx, req.GetTripId())
	if err != nil {
		if errors.Is(err, service.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "can't get trip: %v", err)
	}

	return &driver.GetTripByIDResponse{
		Trip: converters.TripFromEntity(trip),
	}, nil
}
