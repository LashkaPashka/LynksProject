package db

import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)
type MysqlDb struct {
	Db *sql.DB
}


func New(connstr string) (*MysqlDb, error) {
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		return nil, err
	}
	
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MysqlDb{
		Db: db,
	}, nil
}