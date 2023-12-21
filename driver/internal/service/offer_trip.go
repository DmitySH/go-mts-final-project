package service

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (d *DriverService) OfferTripForDrivers(ctx context.Context, offerID, tripID string) error {
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("offerID", offerID)))
	defer span.End()

	offer, err := d.offeringClient.GetOfferByID(ctx, offerID)
	if err != nil {
		return fmt.Errorf("can't get trip: %w", err)
	}

	err = d.repo.CreateTrip(ctx, entity.Trip{
		Id: tripID,
		From: entity.LatLng{
			Lat: offer.From.Lat,
			Lng: offer.From.Lng,
		},
		To: entity.LatLng{
			Lat: offer.To.Lat,
			Lng: offer.To.Lng,
		},
		Price: entity.Money{
			Amount:   offer.Price.Amount,
			Currency: offer.Price.Currency,
		},
		Status: entity.TripStatusDriverSearch,
	})

	if err != nil {
		return fmt.Errorf("can't create trip: %w", err)
	}

	drivers, err := d.locationClient.GetDrivers(ctx, offer.From, d.cfg.OfferRadius)
	if err != nil {
		return fmt.Errorf("can't get drivers: %w", err)
	}

	for _, driver := range drivers {
		err = d.notifier.NotifyDriver(ctx, driver.Id, offer)
		if err != nil {
			loggy.Errorln("can't notify driver:", err)
		}
	}

	return nil
}
