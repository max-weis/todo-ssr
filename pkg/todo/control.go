package todo

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
)

var (
	ValidationError error = errors.New("validation error")

	ErrNoText      = fmt.Errorf("%w: text is required", ValidationError)
	ErrTextTooLong = fmt.Errorf("%w: text must be less than 256 characters", ValidationError)
)

const MAX_TEXT int = 256

type (
	// Todo describes a todo item.
	Todo struct {
		// Text is the content of the todo item. It must be unique.
		Text string
		// Done indicates whether the todo item is done.
		Done bool
	}

	// Page is used to paginate todo items.
	Page struct {
		// Limit is the maximum number of todo items to return.
		Limit int
		// Offset is the number of todo items to skip.
		Offset int
	}

	// Controller provides todo-related operations.
	Controller struct {
		logger     *slog.Logger
		Repository Repository
	}

	// Repository provides access to the todo database.
	Repository interface {
		// Save persists a todo item.
		Save(ctx context.Context, todo Todo) error
		// List returns a list of todo items.
		List(ctx context.Context, page Page) ([]Todo, error)
		// Update updates a todo item.
		Update(ctx context.Context, old, new Todo) error
	}
)

// NewController creates a new todo controller.
func NewController(logger *slog.Logger, repo Repository) Controller {
	return Controller{logger: logger, Repository: repo}
}

// CreateNewTodo creates a new todo item.
func (c *Controller) CreateNewTodo(ctx context.Context, todo Todo) error {
	if err := validateTodo(todo); err != nil {
		c.logger.Error("failed to validate todo", slog.String("text", todo.Text), slog.String("error", err.Error()))
		return err
	}
	if err := c.Repository.Save(ctx, todo); err != nil {
		c.logger.Error("failed to save todo", slog.String("text", todo.Text), slog.String("error", err.Error()))
		return err
	}

	c.logger.Info("todo saved", slog.String("text", todo.Text))

	return nil
}

// ListTodos returns a list of todo items.
func (c *Controller) ListTodos(ctx context.Context, page Page) ([]Todo, error) {
	todos, err := c.Repository.List(ctx, page)
	if err != nil {
		c.logger.Error("failed to list todos", slog.String("error", err.Error()))
		return nil, err
	}

	c.logger.Info("todos listed", slog.Int("count", len(todos)))

	return todos, nil
}

// UpdateTodo updates a todo item.
func (c *Controller) UpdateTodo(ctx context.Context, old, new Todo) error {
	if old == new {
		c.logger.Info("todo not updated", slog.String("text", old.Text))
		return nil
	}

	if err := validateTodo(new); err != nil {
		c.logger.Error("failed to validate todo", slog.String("text", old.Text), slog.String("error", err.Error()))
		return err
	}

	if err := c.Repository.Update(ctx, old, new); err != nil {
		c.logger.Error("failed to update todo", slog.String("text", old.Text), slog.String("error", err.Error()))
		return err
	}

	c.logger.Info("todo updated", slog.String("text", old.Text))

	return nil
}

// validateTodo validates a todo item.
func validateTodo(todo Todo) error {
	if todo.Text == "" {
		return ErrNoText
	}

	if len(todo.Text) > MAX_TEXT {
		return ErrTextTooLong
	}

	return nil
}
