package server

import (
	"Lynks/user/configs"
	"Lynks/user/internal/api"
	"Lynks/user/internal/db"
	"Lynks/user/pkg/logger"
	"log/slog"
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

func (s *Server) NewLogger(){
	log := logger.SetupLogger(s.conf.Env)

	log.Info("Starting application",
			slog.String("Addr", "127.0.0.1"),
			slog.String("Port", "8082"),
			slog.String("Env", s.conf.Env),
		)
	
	log.Info("Db's running",
			slog.String("Port", "5432"),
		)
}