package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	Env string
	DSN DSNConfig
	Secret string
}

type DSNConfig struct {
	DSN string
}

func LoadConfig() (*Configs, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	
	return &Configs{
		Env: os.Getenv("ENV"),
		DSN: DSNConfig{
			DSN: os.Getenv("DSN"),
		},
		Secret: os.Getenv("SECRET"),
	}, nil

}