package service

import (
	"context"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
)

type DriverConfig struct {
	OfferRadius float64
}

type DriverRepository interface {
	GetTrips(ctx context.Context) ([]entity.Trip, error)
	GetTripByID(ctx context.Context, tripID string) (entity.Trip, error)
	UpdateTripStatus(ctx context.Context, tripID, driverID string, tripStatus entity.TripStatus) error
	UpdateTripDriver(ctx context.Context, tripID string, driverID string) error
	CreateTrip(ctx context.Context, trip entity.Trip) error
}

type DriverProducer interface {
	ProduceJSONMessage(ctx context.Context, data any) error
}

type LocationClient interface {
	GetDrivers(ctx context.Context, location entity.LatLng, radius float64) ([]entity.Driver, error)
}

type OfferingClient interface {
	GetOfferByID(ctx context.Context, offerID string) (entity.Offer, error)
}

type DriverNotifier interface {
	NotifyDriver(ctx context.Context, driverID string, offer entity.Offer) error
}

type DriverService struct {
	cfg DriverConfig

	repo     DriverRepository
	producer DriverProducer
	notifier DriverNotifier

	locationClient LocationClient
	offeringClient OfferingClient
}

func NewDriverService(cfg DriverConfig, repo DriverRepository, producer DriverProducer, notifier DriverNotifier,
	locationClient LocationClient, offeringClient OfferingClient) *DriverService {
	return &DriverService{
		cfg:            cfg,
		repo:           repo,
		producer:       producer,
		notifier:       notifier,
		locationClient: locationClient,
		offeringClient: offeringClient,
	}
}
