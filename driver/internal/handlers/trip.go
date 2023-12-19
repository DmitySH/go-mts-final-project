package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"time"
)

type messageType string

const (
	createdTripEvent = messageType("trip.event.created")
)

var (
	ErrUnknownMessageType = errors.New("unknown message type")
)

type CreatedTripEvent struct {
	Id              string    `json:"id"`
	Source          string    `json:"source"`
	Type            string    `json:"type"`
	Datacontenttype string    `json:"datacontenttype"`
	Time            time.Time `json:"time"`
	Data            struct {
		TripId  string `json:"trip_id"`
		OfferId string `json:"offer_id"`
	} `json:"data"`
}

type TripHandler struct {
}

func NewTripHandler() *TripHandler {
	return &TripHandler{}
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
	var event CreatedTripEvent

	err := json.Unmarshal(m.Value, &event)
	if err != nil {
		return err
	}

	loggy.Infoln(event)

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
