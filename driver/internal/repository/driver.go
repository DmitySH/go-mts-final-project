package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName        = "driver"
	tripsCollectionName = "trips"
)

type DriverRepository struct {
	tripsColl *mongo.Collection
}

func NewDriverRepository(db *mongo.Client) *DriverRepository {
	return &DriverRepository{
		tripsColl: db.Database(databaseName).Collection(tripsCollectionName),
	}
}
