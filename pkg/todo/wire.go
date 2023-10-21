package todo

import "github.com/google/wire"

var Providers = wire.NewSet(
	NewSqliteRepository,
	NewController,
)
