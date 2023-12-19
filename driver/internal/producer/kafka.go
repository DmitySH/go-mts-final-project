package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
)

type KafkaProducer struct {
	w *kafka.Writer
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func NewProducer(cfg KafkaConfig) (*KafkaProducer, func() error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Brokers...),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{w: w}, w.Close
}

func (p *KafkaProducer) ProduceMessage(ctx context.Context, msg entity.QMessage) error {
	err := p.w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(msg.Key),
			Value: msg.Value,
		},
	)

	return err
}
