package service

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (l *LocationService) GetDrivers(ctx context.Context, latLng entity.LatLng, radius float64) ([]entity.Driver, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	return l.repo.GetDriversInsideCircle(ctx, latLng, radius)
}
