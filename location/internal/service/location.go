package service

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/entity"
)

type LocationRepository interface {
	UpdateDriverLocation(ctx context.Context, driverID uuid.UUID, latLon entity.LatLng) error
	GetDriversInsideCircle(ctx context.Context, latLon entity.LatLng, radius float64) ([]entity.Driver, error)
}

type LocationService struct {
	repo LocationRepository
}

func NewLocationService(repo LocationRepository) *LocationService {
	return &LocationService{
		repo: repo,
	}
}
