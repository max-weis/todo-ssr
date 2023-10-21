package todo_test

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/max-weis/todo-ssr/internal/database"

	"github.com/max-weis/todo-ssr/pkg/todo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SQLiteRepositoryTestSuite struct {
	suite.Suite
	ctx        context.Context
	repository todo.Repository
	db         *sqlx.DB
}

func (suite *SQLiteRepositoryTestSuite) SetupTest() {
	db, err := sqlx.Open("sqlite3", ":memory:")
	suite.NoError(err)

	err = database.Migrate(db)
	suite.NoError(err)

	db.MustExec("INSERT INTO todos (text, done) VALUES ('clean kitchen', false); " +
		"INSERT INTO todos (text, done) VALUES ('clean bathroom', true); " +
		"INSERT INTO todos (text, done) VALUES ('clean bedroom', false);")

	suite.db = db
	suite.repository = todo.NewSqliteRepository(db)
	suite.ctx = context.Background()
}

func (suite *SQLiteRepositoryTestSuite) TestRepository_Save() {
	err := suite.repository.Save(suite.ctx, todo.Todo{
		Text: "clean kitchen",
		Done: false,
	})

	suite.NoError(err)
}

func (suite *SQLiteRepositoryTestSuite) TestRepository_List() {
	todos, err := suite.repository.List(suite.ctx, todo.Page{
		Limit:  10,
		Offset: 0,
	})

	suite.Len(todos, 3)
	suite.NoError(err)
}

func (suite *SQLiteRepositoryTestSuite) TestRepository_Update() {
	oldTodo := todo.Todo{
		Text: "clean kitchen",
		Done: false,
	}

	newTodo := todo.Todo{
		Text: "clean kitchen",
		Done: true,
	}

	err := suite.repository.Update(suite.ctx, oldTodo, newTodo)
	suite.NoError(err)

	var updatedDone bool
	err = suite.db.Get(&updatedDone, "SELECT done FROM todos WHERE text = 'clean kitchen' LIMIT 1")
	suite.NoError(err)

	suite.True(updatedDone)
}

func TestSQLiteRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SQLiteRepositoryTestSuite))
}
