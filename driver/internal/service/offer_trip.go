package service

import (
	"context"
	"fmt"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (d *DriverService) OfferTripForDrivers(ctx context.Context, offerID string) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()
	
	offer, err := d.offeringClient.GetOfferByID(ctx, offerID)
	if err != nil {
		return fmt.Errorf("can't get trip: %w", err)
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
