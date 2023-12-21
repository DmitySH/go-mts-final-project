package service

import (
	"context"
	"fmt"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/messages"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

func (d *DriverService) AcceptTrip(ctx context.Context, tripID string) error {
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("tripID", tripID)))
	defer span.End()

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
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("tripID", tripID)))
	defer span.End()

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
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("tripID", tripID)))
	defer span.End()

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
