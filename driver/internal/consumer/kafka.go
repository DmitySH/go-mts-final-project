package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/loggy"
	"go.uber.org/multierr"
	"io"
)

type Consumer struct {
	cfg KafkaConfig

	topicHandlers map[string]*topicHandler
}

type topicHandler struct {
	topic   string
	reader  *kafka.Reader
	handler Handler

	stopCtx context.Context
}

func NewConsumer(cfg KafkaConfig, opts ...ConsumerOption) *Consumer {
	c := &Consumer{
		cfg:           cfg,
		topicHandlers: make(map[string]*topicHandler),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Consumer) Run(ctx context.Context) {
	for topic, th := range c.topicHandlers {
		loggy.Infoln("starting consuming messages for topic", topic)
		go c.readMessages(ctx, th)
	}
}

func (c *Consumer) readMessages(ctx context.Context, th *topicHandler) {
	for {
		m, err := th.reader.ReadMessage(ctx)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			loggy.Errorf("can't read message from topic %s: %v", th.topic, err)
			continue
		}

		err = th.handler.Handle(ctx, m)
		if err != nil {
			loggy.Errorf("can't handle message from topic %s: %v", th.topic, err)
			continue
		}
	}
}

func (c *Consumer) Stop() error {
	var err error
	for topic, th := range c.topicHandlers {
		err = th.reader.Close()
		if err != nil {
			err = multierr.Append(err, fmt.Errorf("can't close kafka reader for topic %s: %w", topic, err))
		}
	}

	return err
}
