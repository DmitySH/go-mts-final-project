package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository struct {
	db *mongo.Client
}

func NewDriverRepository(db *mongo.Client) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}
