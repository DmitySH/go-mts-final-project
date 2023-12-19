package service

import (
	"context"
	"fmt"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
)

func (d *DriverService) OfferTripForDrivers(ctx context.Context, offerID string) error {
	trip, err := d.offeringClient.GetOfferByID(ctx, offerID)
	if err != nil {
		return fmt.Errorf("can't get trip: %w", err)
	}

	drivers, err := d.locationClient.GetDrivers(ctx, trip.From, d.cfg.OfferRadius)
	if err != nil {
		return fmt.Errorf("can't get drivers: %w", err)
	}

	for _, driver := range drivers {
		// TODO: Offer trip via websocket
		loggy.Infoln(driver)
	}

	return nil
}
