package repository

import (
	"context"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/entity"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/location/internal/service"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

func (l *LocationRepository) GetDriversInsideCircle(ctx context.Context, latLon entity.LatLng, radius float64) ([]entity.Driver, error) {
	ctx, span := tracy.Start(ctx)
	defer span.End()

	sqlQuery, args, err := l.psql.Select("driver.id, driver.name AS name, auto.name AS auto").
		From(driverTable).
		Where("(lat - ?) * (lat - ?) + (lng - ?) * (lng - ?) <= ?::DOUBLE PRECISION * ?", latLon.Lat, latLon.Lat, latLon.Lng, latLon.Lng, radius, radius).
		Join(carTable + " ON auto.id = auto").
		ToSql()
	if err != nil {
		return nil, buildSQLError(ctx, err)
	}

	var drivers []entity.Driver
	err = l.db.SelectContext(ctx, &drivers, sqlQuery, args...)
	if err != nil {
		return nil, executeSQLError(ctx, err)
	}

	if len(drivers) == 0 {
		return nil, service.ErrEntityNotFound
	}

	return drivers, nil
}
