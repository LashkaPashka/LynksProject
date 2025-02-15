package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	Db DSNconfig
}

type DSNconfig struct {
	DSN string
}

func LoadConfig() (*Config, error){
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	
	return &Config{
		Env: os.Getenv("ENV"),
		Db: DSNconfig{
			DSN: os.Getenv("DSN"),
		},
	}, nil
}