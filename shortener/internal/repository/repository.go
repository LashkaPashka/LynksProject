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
		)
	}

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


func (repo *LinkRepository) GetLinks(ctx context.Context) ([]model.Links, error) {
	rows, err := repo.db.Query(ctx, `SELECT id, url, hash FROM links`)
	if err != nil {
		return nil, err
	}

	var links []model.Links
	for rows.Next() {
		var l model.Links
		err := rows.Scan(
			&l.ID,
			&l.Url,
			&l.ShortUrl,
		)
		if err != nil {
			return nil, err
		}

		links = append(links, l)
	}

	if err := rows.Err(); err != nil {
		return nil, err 
	}

	return links, nil
}


func (repo *LinkRepository) CreateLinks(ctx context.Context, links []model.Links) error {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var batch = &pgx.Batch{}
	

	for _, link := range links {
		batch.Queue(`INSERT INTO links(url, hash) VALUES($1, $2)`, link.Url, link.ShortUrl)
	} 

	res := tx.SendBatch(ctx, batch)
	if err := res.Close(); err != nil {
		return err
	}

	return tx.Commit(ctx)
}