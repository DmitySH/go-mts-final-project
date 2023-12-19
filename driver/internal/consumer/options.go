package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type Handler interface {
	Handle(ctx context.Context, m kafka.Message) error
}

type ConsumerOption func(c *Consumer)

func WithHandler(topic string, handler Handler) ConsumerOption {
	return func(c *Consumer) {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers: c.cfg.Brokers,
			GroupID: c.cfg.GroupID,
			Topic:   topic,
			MaxWait: time.Second * 3,
		})

		c.topicHandlers[topic] = &topicHandler{
			topic:   topic,
			reader:  r,
			handler: handler,
		}
	}
}
