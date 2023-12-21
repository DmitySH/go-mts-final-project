package server

import (
	"context"
	"fmt"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/pkg/api/driver"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (d *DriverServer) AcceptTrip(ctx context.Context, req *driver.AcceptTripRequest) (*driver.AcceptTripResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	err := d.driverService.AcceptTrip(ctx, req.GetTripId())
	if err != nil {
		return nil, fmt.Errorf("can't accept trip: %w", err)
	}

	return &driver.AcceptTripResponse{}, nil
}
func (d *DriverServer) StartTrip(ctx context.Context, req *driver.AcceptTripRequest) (*driver.AcceptTripResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	err := d.driverService.StartTrip(ctx, req.GetTripId())
	if err != nil {
		return nil, fmt.Errorf("can't start trip: %w", err)
	}

	return &driver.AcceptTripResponse{}, nil
}

func (d *DriverServer) CancelTrip(ctx context.Context, req *driver.CancelTripRequest) (*driver.CancelTripResponse, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	err := d.driverService.CancelTrip(ctx, req.GetTripId())
	if err != nil {
		return nil, fmt.Errorf("can't cancel trip: %w", err)
	}

	return &driver.CancelTripResponse{}, nil
}
