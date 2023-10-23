package todo

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type (
	sqliteRepository struct {
		db *sqlx.DB
	}

	todoEntity struct {
		Text string `db:"text"`
		Done bool   `db:"done"`
	}
)

// NewSqliteRepository returns an SQLite implementation of the Repository interface
func NewSqliteRepository(db *sqlx.DB) Repository {
	return sqliteRepository{db: db}
}

func (s sqliteRepository) Save(ctx context.Context, todo Todo) error {
	const query = `INSERT INTO todos (text, done) VALUES (:text, :done)`

	_, err := s.db.NamedExecContext(ctx, query, mapToEntity(todo))
	return err
}

func (s sqliteRepository) List(ctx context.Context) ([]Todo, error) {
	const query = `SELECT text, done FROM todos`

	var entities []todoEntity
	if err := s.db.SelectContext(ctx, &entities, query); err != nil {
		return nil, err
	}

	todos := make([]Todo, len(entities))
	for i, entity := range entities {
		todos[i] = mapFromEntity(entity)
	}

	return todos, nil
}

func (s sqliteRepository) Update(ctx context.Context, old, new Todo) error {
	const query = `UPDATE todos SET text = :new_text, done = :new_done WHERE text = :old_text`

	_, err := s.db.NamedExecContext(ctx, query, map[string]interface{}{
		"old_text": old.Text,
		"new_text": new.Text,
		"new_done": new.Done,
	})
	return err
}

func mapToEntity(todo Todo) todoEntity {
	return todoEntity{
		Text: todo.Text,
		Done: todo.Done,
	}
}

func mapFromEntity(entity todoEntity) Todo {
	return Todo{
		Text: entity.Text,
		Done: entity.Done,
	}
}
