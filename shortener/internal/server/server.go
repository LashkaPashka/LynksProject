package server

import (
	"Lynks/shortener/internal/api"
	"Lynks/shortener/internal/db"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	api *api.API
	db *db.Db
}

func New() (*Server, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	connstr := os.Getenv("DSN")
	db, err := db.New(connstr)
	if err != nil {
		return nil, err
	}

	s := &Server{}

	s.db = db
	s.api = api.New(s.db)

	return s, nil
}

func (s *Server) Run(){
	fmt.Println("Server's listening 8081")
	s.api.Run(":8081")
}