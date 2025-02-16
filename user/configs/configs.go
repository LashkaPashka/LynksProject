package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	env string
	DSN DSNConfig
}

type DSNConfig struct {
	DSN string
}

func LoadConfig() (*Configs, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	
	return &Configs{
		env: os.Getenv("ENV"),
		DSN: DSNConfig{
			DSN: os.Getenv("DSN"),
		},
	}, nil

}