package server

import (
	"memCache/internal/api"
	"memCache/internal/db"
)

type Server struct {
	db *db.Db
	api *api.API
}

func New() (*Server, error) {
	server := &Server{}
	
	db, err := db.New()
	if err != nil {
		return nil, err
	}
	
	server.db = db
	server.api, err = api.New(server.db)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (s *Server) Run() {
	s.api.Run(":8084")
}