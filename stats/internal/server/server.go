package server

import (
	"Stats/configs"
	"Stats/internal/db"
	"Stats/pkg/logger"
	"Stats/pkg/queue"
	"log/slog"
	"net/http"
)

type Server struct {
	db *db.MysqlDb
	Conf *configs.Config
	kafka *kafka.Client
}

func New() (*Server, error) {
	conf, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	
	server := &Server{
		Conf: conf,	
	}
	
	kafka, err := kafka.New([]string{"kafka"}, "test-topic", "test-consumer-group", server.Conf)
	if err != nil {
		return nil, err
	}
	db, err := db.New(conf.Db.DSN)
	if err != nil {
		return nil, err
	}

	server.kafka = kafka
	server.db = db
	

	return server, nil
}

func (s *Server) Run() {
	http.ListenAndServe(":8083", nil)
}

func (s *Server) NewLogger() {
		log := logger.SetupLogger(s.Conf.Env)
	
		log.Info("Starting application STATS",
				slog.String("Addr", "127.0.0.1"),
				slog.String("Port", "8083"),
				slog.String("Env", s.Conf.Env),
			)
		
		log.Info("Db's running",
				slog.String("Port", "3306"),
				slog.String("DB_NAME", "mySQL"),
			)
}

func (s *Server) KafkaConsumer() {
	go s.kafka.Consumer()
}