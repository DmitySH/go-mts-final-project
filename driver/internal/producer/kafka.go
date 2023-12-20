package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
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
	payload, err := json.Marshal(&data)
	if err != nil {
		return fmt.Errorf("can't marshal data: %w", err)
	}

	err = p.w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(uuid.New().String()),
		Value: payload,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProducer) Close() error {
	return p.w.Close()
}
