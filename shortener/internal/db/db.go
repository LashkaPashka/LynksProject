package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	pool *pgxpool.Pool
}

func New(connstr string) (*Db, error) {
	pool, err := pgxpool.New(context.Background(), connstr)

	if err != nil {
		return nil, err
	}
	
	return &Db{
		pool: pool,
	}, nil
}