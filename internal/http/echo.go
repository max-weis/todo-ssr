package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/max-weis/todo-ssr/internal/config"
	"log/slog"
)

func NewEcho(cfg *config.Config, logger *slog.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod: true,
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("received request", slog.String("method", v.Method), slog.String("uri", v.URI), slog.Int("status", v.Status))
			return nil
		},
	}))

	return e
}
