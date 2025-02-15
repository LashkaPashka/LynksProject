package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev = "Dev"
	envProd = "Prod"
)

var (
	Log *slog.Logger
)

func SetupLogger(env string) *slog.Logger {
	switch env {
		case envLocal:
			Log = slog.New(
				slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		case envDev:
			Log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		case envProd:
			Log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
	}

	return Log
}