package service

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
)

type DriverRepository interface {
}

type DriverProducer interface {
	ProduceMessage(ctx context.Context, msg entity.QMessage) error
}

type DriverService struct {
	repo     DriverRepository
	producer DriverProducer
}

func NewDriverService(repo DriverRepository, producer DriverProducer) *DriverService {
	return &DriverService{
		repo:     repo,
		producer: producer,
	}
}
