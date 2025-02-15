package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	pool *pgxpool.Pool
}

var ctx = context.Background()

func New(connstr string) (*Db, error) {
	pool, err := pgxpool.New(ctx, connstr)
	if err != nil {
		return nil, err
	}
	
	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Db{
		pool: pool,
	}, nil
}