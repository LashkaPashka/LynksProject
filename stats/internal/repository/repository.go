package repository

import (
	"Stats/internal/model"
	"context"
	"database/sql"
	"fmt"

	"Stats/configs"

	_ "github.com/go-sql-driver/mysql"
)

type StatsRepository struct {
	db *sql.DB
	conf *configs.Config
}

func NewStatRepostiory(conf *configs.Config) (*StatsRepository, error) {
	db, err := sql.Open("mysql", conf.Db.DSN)
	if err != nil {
		return nil, err
	}
	
	stat := &StatsRepository{
		db: db,
		conf: conf,
	}
	
	return stat, nil
}
 

func (repo *StatsRepository) GetStatByUrl(url string) (*model.Stats, bool, error) {
	rows, err := repo.db.Query(`SELECT url, click, average_length FROM stats WHERE url = ?`, url)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	var stat model.Stats
	
	for rows.Next() {
		err := rows.Scan(
			&stat.Url,
			&stat.Clicks,
			&stat.Average_length,
		)
		if err != nil {
			return nil, false, err
		}
	}
	
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}

	isExistStat := stat != model.Stats{}

	return &stat, isExistStat, nil
}


func (repo *StatsRepository) CreateStat(stats *model.Stats) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO stats(url, click, average_length) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(context.Background(), stats.Url, stats.Clicks, stats.Average_length)
	if err != nil {
		return err
	}


	id, _ := res.LastInsertId()
	fmt.Println("Запись элемента", id)

	return tx.Commit()
}

func (repo *StatsRepository) UpdateStat(stat *model.Stats) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`UPDATE stats SET click = ? WHERE url = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.TODO(), stat.Clicks, stat.Url)
	if err != nil {
		return err
	}

	return tx.Commit()
}