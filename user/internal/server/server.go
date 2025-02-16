package server

import (
	"Lynks/user/configs"
	"Lynks/user/internal/api"
	"Lynks/user/internal/db"
)

type Server struct {
	mongoDb *db.MongoDb
	api *api.API
	conf *configs.Configs
}

func New() (*Server, error) {
	// Loading configs
	conf, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	// Settings Server
	
	server := &Server{}
	
	mongoDb, err := db.New(conf.DSN.DSN)
	if err != nil {
		return nil, err
	}

	server.mongoDb = mongoDb
	server.api = api.New(server.mongoDb)
	server.conf = conf	

	return server, nil
}

func (s *Server) Run() {
	s.api.Run(":8082")
}