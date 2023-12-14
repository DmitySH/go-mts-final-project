package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"time"
)

const (
	defaultPingTimeout = 3 * time.Second
)

type MongoConfig struct {
	URI        string
	Username   string
	Password   string
	AuthSource string
}

func NewMongoDB(ctx context.Context, cfg MongoConfig) (*mongo.Client, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.URI)
	optsAuth := options.Credential{
		Username:   cfg.Username,
		Password:   cfg.Password,
		AuthSource: cfg.AuthSource,
	}

	opts.SetAuth(optsAuth)
	opts.Monitor = otelmongo.NewMonitor()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongo: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, defaultPingTimeout)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("can't ping mongo: %w", err)
	}

	return client, nil
}
