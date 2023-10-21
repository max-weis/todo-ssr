//go:build wireinject

//go:generate wire .
package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/max-weis/todo-ssr/internal"
	"github.com/max-weis/todo-ssr/internal/config"
	"github.com/max-weis/todo-ssr/pkg/todo"
	"log/slog"
)

type App struct {
	*echo.Echo
	Log     *slog.Logger
	cfg     *config.Config
	handler *todo.Handler
}

func InitApp() (*App, error) {
	panic(wire.Build(
		internal.Providers,
		todo.Providers,

		wire.Struct(new(App), "*"),
	))
}
