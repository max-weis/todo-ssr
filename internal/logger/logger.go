package logger

import (
	"github.com/max-weis/todo-ssr/internal/config"
	"log/slog"
	"os"
)

func NewLogger(cfg *config.Config) *slog.Logger {
	var logger *slog.Logger

	if !cfg.Production {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	return logger
}
