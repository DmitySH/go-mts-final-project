package service

import (
	"context"
	"fmt"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/messages"
	"time"
)

func (d *DriverService) AcceptTrip(ctx context.Context, tripID string) error {
	// TODO: get trip, convert to messages.Command, marshal
	cmd := messages.AcceptedTripCommand{
		Id:              "",
		Source:          "",
		Type:            "",
		Datacontenttype: "",
		Time:            time.Time{},
		Data: struct {
			TripId string `json:"trip_id"`
		}{},
	}

	err := d.producer.ProduceJSONMessage(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't produce message: %w", err)
	}

	return nil
}

func (d *DriverService) StartTrip(ctx context.Context, tripID string) error {
	// TODO: get trip, convert to messages.Command, marshal
	cmd := messages.StartedTripCommand{
		Id:              "",
		Source:          "",
		Type:            "",
		Datacontenttype: "",
		Time:            time.Time{},
		Data: struct {
			TripId string `json:"trip_id"`
		}{},
	}

	err := d.producer.ProduceJSONMessage(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't produce message: %w", err)
	}

	return nil
}

func (d *DriverService) CancelTrip(ctx context.Context, tripID string) error {
	// TODO: get trip, convert to messages.Command, marshal
	cmd := messages.CancelledTripCommand{
		Id:              "",
		Source:          "",
		Type:            "",
		Datacontenttype: "",
		Time:            time.Time{},
		Data: struct {
			TripId string `json:"trip_id"`
		}{},
	}

	err := d.producer.ProduceJSONMessage(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't produce message: %w", err)
	}

	return nil
}
