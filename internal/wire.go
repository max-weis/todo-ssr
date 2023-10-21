package internal

import (
	"github.com/google/wire"
	"github.com/max-weis/todo-ssr/internal/config"
	"github.com/max-weis/todo-ssr/internal/database"
	"github.com/max-weis/todo-ssr/internal/http"
	"github.com/max-weis/todo-ssr/internal/logger"
)

var Providers = wire.NewSet(
	config.NewConfig,
	logger.NewLogger,
	http.NewEcho,
	database.NewDatabase,
)
