package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/messages"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
)

type messageType string

const (
	createdTripEvent = messageType("trip.event.created")
)

var (
	ErrUnknownMessageType = errors.New("unknown message type")
)

type TripHandler struct {
	driverService *service.DriverService
}

func NewTripHandler(driverService *service.DriverService) *TripHandler {
	return &TripHandler{
		driverService: driverService,
	}
}

func (c *TripHandler) Handle(ctx context.Context, m kafka.Message) error {
	msgType, err := getTypeFromMessage(m)
	if err != nil {
		return fmt.Errorf("can't get type from message: %w", err)
	}

	switch msgType {
	case createdTripEvent:
		err = c.handleCreatedTripEvent(ctx, m)
	default:
		err = fmt.Errorf("%s: %w", msgType, ErrUnknownMessageType)
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *TripHandler) handleCreatedTripEvent(ctx context.Context, m kafka.Message) error {
	var event messages.CreatedTripEvent

	err := json.Unmarshal(m.Value, &event)
	if err != nil {
		return fmt.Errorf("can't unmarshal event: %w", err)
	}

	err = c.driverService.OfferTripForDrivers(ctx, event.Data.OfferId)
	if err != nil {
		return fmt.Errorf("can't offer trip: %w", err)
	}

	return nil
}

func getTypeFromMessage(m kafka.Message) (messageType, error) {
	type typeContainer struct {
		Type string `json:"type"`
	}

	t := typeContainer{}

	err := json.Unmarshal(m.Value, &t)
	if err != nil {
		return "", err
	}

	return messageType(t.Type), nil
}
