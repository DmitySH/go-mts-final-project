package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

type KafkaProducer struct {
	w *kafka.Writer
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func NewProducer(cfg KafkaConfig) *KafkaProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Brokers...),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{w: w}
}

func (p *KafkaProducer) ProduceJSONMessage(ctx context.Context, data any) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	payload, err := json.Marshal(&data)
	if err != nil {
		err = fmt.Errorf("can't marshal data: %w", err)
		tracy.RecordError(ctx, err)
		return err
	}

	err = p.w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(uuid.New().String()),
		Value: payload,
	})
	if err != nil {
		tracy.RecordError(ctx, err)
		return err
	}

	return nil
}

func (p *KafkaProducer) Close() error {
	return p.w.Close()
}
