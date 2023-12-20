package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
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

func createCursorError(ctx context.Context, err error) error {
	err = fmt.Errorf("can't create cursor: %w", err)
	tracy.RecordError(ctx, err)
	return err
}

func executeError(ctx context.Context, err error) error {
	err = fmt.Errorf("can't execute query: %w", err)
	tracy.RecordError(ctx, err)
	return err
}

func (d *DriverRepository) GetTrips(ctx context.Context) ([]entity.Trip, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	cursor, err := d.tripsColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, createCursorError(ctx, err)
	}

	defer cursor.Close(ctx)

	var trips []entity.Trip
	if err := cursor.All(ctx, trips); err != nil {
		return nil, executeError(ctx, err)
	}

	return trips, nil
}

func (d *DriverRepository) GetTripById(ctx context.Context, tripId string) (entity.Trip, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	cursor, err := d.tripsColl.Find(ctx, bson.M{"id": tripId})
	if err != nil {
		return entity.Trip{}, createCursorError(ctx, err)
	}

	defer cursor.Close(ctx)

	var trips []entity.Trip
	if err := cursor.All(ctx, trips); err != nil {
		return entity.Trip{}, executeError(ctx, err)
	}

	if len(trips) != 1 {
		return entity.Trip{}, errors.New("number of trips is not 1")
	}

	return trips[0], nil
}

func (d *DriverRepository) UpdateTripStatus(ctx context.Context, tripId string, tripStatus entity.TripStatus) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	if _, err := d.tripsColl.UpdateOne(ctx, bson.M{"id": tripId}, bson.M{"$set": bson.M{"status": tripStatus}}); err != nil {
		return executeError(ctx, err)
	}

	return nil
}

func (d *DriverRepository) UpdateTripDriver(ctx context.Context, tripId string, driverId string) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	if _, err := d.tripsColl.UpdateOne(ctx, bson.M{"id": tripId}, bson.M{"$set": bson.M{"driver": driverId}}); err != nil {
		return executeError(ctx, err)
	}

	return nil
}

func (d *DriverRepository) CreateTrip(ctx context.Context, trip entity.Trip) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	cursor, err := d.tripsColl.Find(ctx, bson.M{"id": trip.Id})
	if err != nil {
		return createCursorError(ctx, err)
	}

	var trips []entity.Trip
	if err := cursor.All(ctx, trips); err != nil {
		return executeError(ctx, err)
	}

	defer cursor.Close(ctx)

	if len(trips) != 0 {
		return errors.New("trip already exist")
	}

	_, err = d.tripsColl.InsertOne(ctx, trip)
	return err
}
