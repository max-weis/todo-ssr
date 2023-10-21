package database_test

import (
	"github.com/max-weis/todo-ssr/internal/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db, err := database.NewDatabase()
	assert.NoError(t, err)

	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='todos';").Scan(&tableName)
	assert.NoError(t, err)

	assert.Equal(t, "todos", tableName)
}
