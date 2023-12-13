package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (l *LocationRepository) UpdateDriverLocation(ctx context.Context, driverID uuid.UUID, latLon entity.LatLng) error {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	sqlQuery, args, err := l.psql.Update(driverTable).
		Where(sq.Eq{"id": driverID.String()}).
		Set("lat", latLon.Lat).
		Set("lng", latLon.Lng).
		ToSql()
	if err != nil {
		return buildSQLError(ctx, err)
	}

	_, err = l.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return executeSQLError(ctx, err)
	}

	return nil
}
