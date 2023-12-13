package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.com/hse-mts-go-dashagarov/go-taxi/pkg/houston/tracy"
)

const (
	driverTable = "driver"
	carTable    = "auto"
)

type LocationRepository struct {
	db   *sqlx.DB
	psql sq.StatementBuilderType
}

func NewLocationRepository(db *sqlx.DB) *LocationRepository {
	return &LocationRepository{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}
}

func buildSQLError(ctx context.Context, err error) error {
	err = fmt.Errorf("can't build sql: %w", err)
	tracy.RecordError(ctx, err)
	return err
}

func executeSQLError(ctx context.Context, err error) error {
	err = fmt.Errorf("can't execute sql: %w", err)
	tracy.RecordError(ctx, err)
	return err
}
