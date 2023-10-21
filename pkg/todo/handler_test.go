package todo_test

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/max-weis/todo-ssr/internal/database"
	"github.com/max-weis/todo-ssr/pkg/todo"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

type HandlerTestSuite struct {
	suite.Suite
	handler *todo.Handler
	e       *echo.Echo
}

func (suite *HandlerTestSuite) SetupTest() {
	db, err := sqlx.Open("sqlite3", ":memory:")
	suite.NoError(err)

	err = database.Migrate(db)
	suite.NoError(err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	repository := todo.NewSqliteRepository(db)
	controller := todo.NewController(logger, repository)
	e := echo.New()
	handler, err := todo.NewHandler(logger, controller, e)
	suite.NoError(err)
	suite.handler = handler
	suite.e = e
}

func (suite *HandlerTestSuite) TestHandler_Index() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.Index(c)
	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)
}

func (suite *HandlerTestSuite) TestHandler_Add() {
	form := make(url.Values)
	form.Set("todo", "Test Todo")
	req := httptest.NewRequest(http.MethodPost, "/todo", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.Add(c)
	suite.NoError(err)
	suite.Equal(http.StatusCreated, rec.Code)
}

func (suite *HandlerTestSuite) TestHandler_Check() {
	form := make(url.Values)
	form.Set("done", "true")
	form.Set("text", "Test Todo")
	req := httptest.NewRequest(http.MethodPatch, "/todo", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.Check(c)
	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
