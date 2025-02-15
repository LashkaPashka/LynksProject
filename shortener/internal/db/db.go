package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Db struct {
	pool *pgxpool.Pool
}

func New() (*Db, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	connstr := os.Getenv("DSN")

	pool, err := pgxpool.New(context.Background(), connstr)

	if err != nil {
		return nil, err
	}
	
	return &Db{
		pool: pool,
	}, nil
}