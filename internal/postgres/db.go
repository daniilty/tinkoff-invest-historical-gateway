package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ DB = (*DBImpl)(nil)

type DB interface {
	GetCandlesFromTo(context.Context, string, uint32, uint32) (Candles, error)
	InsertCandle(context.Context, Candle) error
}

type DBImpl struct {
	pool *pgxpool.Pool
}

func NewDBImpl(pool *pgxpool.Pool) *DBImpl {
	return &DBImpl{
		pool: pool,
	}
}
