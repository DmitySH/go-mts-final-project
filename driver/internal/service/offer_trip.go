package service

import (
	"context"
	"fmt"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
)

func (d *DriverService) OfferTripForDrivers(ctx context.Context, offerID string) error {
	offer, err := d.offeringClient.GetOfferByID(ctx, offerID)
	if err != nil {
		return fmt.Errorf("can't get trip: %w", err)
	}

	drivers, err := d.locationClient.GetDrivers(ctx, offer.From, d.cfg.OfferRadius)
	if err != nil {
		return fmt.Errorf("can't get drivers: %w", err)
	}

	for i, driver := range drivers {
		loggy.Infoln(i, driver)
		err = d.notifier.NotifyDriver(ctx, driver.Id, offer)
		if err != nil {
			loggy.Errorln("can't notify driver:", err)
		}
	}

	return nil
}
