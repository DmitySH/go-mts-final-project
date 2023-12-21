package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/messages"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (d *DriverService) AcceptTrip(ctx context.Context, tripID, driverID string) error {
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("tripID", tripID),
		attribute.String("driverID", driverID)))
	defer span.End()

	err := d.repo.UpdateTripDriver(ctx, tripID, driverID)
	if err != nil {
		return err
	}

	err = d.repo.UpdateTripStatus(ctx, tripID, driverID, entity.TripStatusDriverFound)
	if err != nil {
		return err
	}

	cmd := messages.AcceptedTripCommand{
		Id:              uuid.New().String(),
		Source:          "/driver",
		Type:            "trip.command.accept",
		Datacontenttype: "application/json",
		Time:            time.Now().UTC(),
		Data: struct {
			TripId   string `json:"trip_id"`
			DriverId string `json:"driver_id"`
		}{
			TripId:   tripID,
			DriverId: driverID,
		},
	}

	err = d.producer.ProduceJSONMessage(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't produce message: %w", err)
	}

	return nil
}

func (d *DriverService) StartTrip(ctx context.Context, tripID, driverID string) error {
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("tripID", tripID),
		attribute.String("driverID", driverID)))
	defer span.End()

	err := d.repo.UpdateTripStatus(ctx, tripID, driverID, entity.TripStatusStarted)
	if err != nil {
		return err
	}

	cmd := messages.StartedTripCommand{
		Id:              uuid.New().String(),
		Source:          "/client",
		Type:            "trip.command.start",
		Datacontenttype: "application/json",
		Time:            time.Now().UTC(),
		Data: struct {
			TripId string `json:"trip_id"`
		}{
			TripId: tripID,
		},
	}

	err = d.producer.ProduceJSONMessage(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't produce message: %w", err)
	}

	return nil
}

func (d *DriverService) CancelTrip(ctx context.Context, tripID, driverID string) error {
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("tripID", tripID),
		attribute.String("driverID", driverID)))
	defer span.End()

	err := d.repo.UpdateTripStatus(ctx, tripID, driverID, entity.TripStatusCanceled)
	if err != nil {
		return err
	}

	cmd := messages.CancelledTripCommand{
		Id:              uuid.New().String(),
		Source:          "/client",
		Type:            "trip.command.cancel",
		Datacontenttype: "application/json",
		Time:            time.Now().UTC(),
		Data: struct {
			TripId string `json:"trip_id"`
		}{
			TripId: tripID,
		},
	}

	err = d.producer.ProduceJSONMessage(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't produce message: %w", err)
	}

	return nil
}

func (d *DriverService) EndTrip(ctx context.Context, tripID, driverID string) error {
	ctx, span := tracy.Start(ctx, trace.WithAttributes(attribute.String("tripID", tripID),
		attribute.String("driverID", driverID)))
	defer span.End()

	err := d.repo.UpdateTripStatus(ctx, tripID, driverID, entity.TripStatusEnded)
	if err != nil {
		return err
	}

	cmd := messages.EndedTripCommand{
		Id:              uuid.New().String(),
		Source:          "/driver",
		Type:            "trip.command.end",
		Datacontenttype: "application/json",
		Time:            time.Now().UTC(),
		Data: struct {
			TripId string `json:"trip_id"`
		}{
			TripId: tripID,
		},
	}

	err = d.producer.ProduceJSONMessage(ctx, cmd)
	if err != nil {
		return fmt.Errorf("can't produce message: %w", err)
	}

	return nil
}
