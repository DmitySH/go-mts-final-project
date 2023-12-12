package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	locationTable = "location"
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
