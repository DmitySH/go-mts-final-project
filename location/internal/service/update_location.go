package service

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (l *LocationService) UpdateDriverLocation(ctx context.Context, driverID uuid.UUID, latLon entity.LatLng) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	return l.repo.UpdateDriverLocation(ctx, driverID, latLon)
}
