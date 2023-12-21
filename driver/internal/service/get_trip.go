package service

import (
	"context"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (d *DriverService) GetTrips(ctx context.Context) ([]entity.Trip, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	return d.repo.GetTrips(ctx)
}

func (d *DriverService) GetTripByID(ctx context.Context, tripID string) (entity.Trip, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	return d.repo.GetTripByID(ctx, tripID)
}
