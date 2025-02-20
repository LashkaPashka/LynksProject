package server

import (
	"ShorteNer/configs"
	"ShorteNer/internal/api"
	"ShorteNer/internal/db"
	"ShorteNer/pkg/logger"
	"log/slog"
)

type Server struct {
	api *api.API
	db *db.Db
	configs *configs.Config
}

func New() (*Server, error){
	configs, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := db.New(configs.Db.DSN)
	if err != nil {
		return nil, err
	}

	s := &Server{}

	s.db = db
	s.configs = configs
	s.api = api.New(s.db, s.configs)	
	
	
	return s, nil
}

func (s *Server) Run(){
	s.api.Run(":8081")
}

func (s *Server) NewLogger(){
	log := logger.SetupLogger(s.configs.Env)

	log.Info("Starting application",
			slog.String("Addr", "127.0.0.1"),
			slog.String("Port", "8081"),
			slog.String("Env", s.configs.Env),
		)
	
	log.Info("Db's running",
			slog.String("Port", "5432"),
			slog.String("DB_NAME", "Postgres"),
		)
}