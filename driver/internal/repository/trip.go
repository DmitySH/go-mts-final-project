package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/driver/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

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

func (d *DriverRepository) GetTripByID(ctx context.Context, tripID string) (entity.Trip, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	cursor, err := d.tripsColl.Find(ctx, bson.M{"id": tripID})
	if err != nil {
		return entity.Trip{}, createCursorError(ctx, err)
	}

	defer cursor.Close(ctx)

	var trips []entity.Trip
	if err := cursor.All(ctx, trips); err != nil {
		return entity.Trip{}, executeError(ctx, err)
	}

	if len(trips) != 1 {
		return entity.Trip{}, service.ErrEntityNotFound
	}

	return trips[0], nil
}

func (d *DriverRepository) UpdateTripStatus(ctx context.Context, tripID, driverID string, tripStatus entity.TripStatus) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	result, err := d.tripsColl.UpdateOne(ctx, bson.M{"id": tripID, "driver_id": driverID}, bson.M{"$set": bson.M{"status": tripStatus}})
	if err != nil {
		return executeError(ctx, err)
	}

	if result.MatchedCount != 1 {
		return service.ErrEntityNotFound
	}

	return nil
}

func (d *DriverRepository) UpdateTripDriver(ctx context.Context, tripID string, driverID string) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	result, err := d.tripsColl.UpdateOne(ctx, bson.M{"id": tripID, "driver_id": bson.M{"$exist": false}}, bson.M{"$set": bson.M{"driver_id": driverID}})
	if err != nil {
		return executeError(ctx, err)
	}

	if result.MatchedCount != 1 {
		return service.ErrEntityNotFound
	}

	return nil
}

func (d *DriverRepository) CreateTrip(ctx context.Context, trip entity.Trip) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	opts := options.Update().SetUpsert(true)
	result, err := d.tripsColl.UpdateOne(ctx, bson.M{"id": trip.Id}, trip, opts)
	if err != nil {
		return executeError(ctx, err)
	}

	if result.MatchedCount != 0 {
		return service.ErrEntityAlreadyExist
	}

	return nil
}
