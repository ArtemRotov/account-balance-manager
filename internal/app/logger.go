package app

import (
	"golang.org/x/exp/slog"
	"os"
)

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

func SetSlog(level string) *slog.Logger {
	var log *slog.Logger
	switch level {
	case local:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case dev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prod:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
