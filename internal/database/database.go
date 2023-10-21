package database

import (
	"embed"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schemas
var schemas embed.FS

func NewDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	if err := Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sqlx.DB) error {
	files, err := schemas.ReadDir("schemas")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := schemas.ReadFile("schemas/" + file.Name())
		if err != nil {
			return err
		}

		if _, err := db.Exec(string(content)); err != nil {
			return err
		}
	}

	return nil
}
