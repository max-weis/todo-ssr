package todo

import (
	"embed"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log/slog"
	"net/http"
)

//go:embed templates/*
var templates embed.FS

type (
	Handler struct {
		logger     *slog.Logger
		controller Controller
		templates  *template.Template
	}

	PageData struct {
		Todos []Todo
	}
)

func NewHandler(logger *slog.Logger, controller Controller, e *echo.Echo) (*Handler, error) {
	tmpl, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		return nil, err
	}

	handler := &Handler{
		logger:     logger,
		controller: controller,
		templates:  tmpl,
	}

	e.Renderer = handler

	e.GET("/", handler.Index)
	e.POST("/todo", handler.Add)
	e.PATCH("/todo", handler.Check)

	return handler, nil
}

func (h *Handler) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if err := h.templates.ExecuteTemplate(w, name, data); err != nil {
		h.logger.Error("failed to render template", slog.String("error", err.Error()))
		return err
	}

	return nil
}

// Index creates the homepage of the application. It lists all the todos.
func (h *Handler) Index(c echo.Context) error {
	todos, err := h.controller.ListTodos(c.Request().Context(), Page{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		return err
	}

	data := PageData{
		Todos: todos,
	}

	return c.Render(http.StatusOK, "index", data)
}

// Add creates a new todo. And returns the new list of todos.
func (h *Handler) Add(c echo.Context) error {
	todoText := c.FormValue("todo")
	todo := Todo{Text: todoText, Done: false}

	if err := h.controller.CreateNewTodo(c.Request().Context(), todo); err != nil {
		return err
	}

	todos, err := h.controller.ListTodos(c.Request().Context(), Page{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		return err
	}

	data := PageData{
		Todos: todos,
	}

	return c.Render(http.StatusCreated, "list", data)
}

// Check toggles the done state of a todo.
func (h *Handler) Check(c echo.Context) error {
	done := c.FormValue("done") == "true"
	text := c.FormValue("text")

	oldTodo := Todo{
		Text: text,
		Done: !done,
	}

	newTodo := Todo{
		Text: text,
		Done: done,
	}

	if err := h.controller.UpdateTodo(c.Request().Context(), oldTodo, newTodo); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
