package repository

import (
	"Lynks/shortener/configs"
	"Lynks/shortener/internal/model"
	"Lynks/shortener/pkg/logger"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)
type myString string
const (
	Hash myString = "hash"
)

type LinkRepository struct {
	conf *configs.Config
	db *pgxpool.Pool
}

func NewLinkRepository() *LinkRepository{
	conf, err := configs.LoadConfig()
	if err != nil {
		logger.Log.Error(
			"Not found env",
			slog.String("Msg", err.Error()),
		) 	}

	repo := LinkRepository{
		conf: conf,
	}
	
	db, err := pgxpool.New(context.Background(), conf.Db.DSN)
	if err != nil {
		logger.Log.Error(
			"Database didn't connect",
			slog.String("Msg", err.Error()),
		)
	}

	repo.db = db
	return &repo
}

func (repo *LinkRepository) GetLinks(ctx context.Context) (string, error) {
	hash := ctx.Value(Hash)
	rows, err := repo.db.Query(ctx, `SELECT url FROM links WHERE hash = $1`, hash.(string))
	if err != nil {
		return "", err
	}
	
	var url string
	
	for rows.Next(){
		err = rows.Scan(
			&url,
		 )
		 if err != nil {
			return "", err
		}
	}

	if err := rows.Err(); err != nil {
		return "", err 
	}
	
	return url, nil
}


func (repo *LinkRepository) CreateLinks(ctx context.Context, link *model.Links) error {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var batch = &pgx.Batch{}
	
	batch.Queue(`INSERT INTO links(url, hash) VALUES($1, $2)`, link.Url, link.Hash)
	
	res := tx.SendBatch(ctx, batch)
	if err := res.Close(); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (repo *LinkRepository) DeleteLinks(ctx context.Context) error {
	hash := ctx.Value(Hash)
	
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var batch = &pgx.Batch{}
	
	batch.Queue(`DELETE FROM links WHERE hash = $1`, hash.(string))
	
	res := tx.SendBatch(ctx, batch)
	if err := res.Close(); err != nil {
		return err
	}

	return tx.Commit(ctx)
}