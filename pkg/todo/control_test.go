//go:generate mockery --name Repository --filename repository_mock_test.go  --output . --outpkg todo_test --structname repositoryMock
package todo_test

// Basic imports
import (
	"context"
	"github.com/max-weis/todo-ssr/pkg/todo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"os"
	"testing"
)

type ControllerTestSuite struct {
	suite.Suite
	ctx        context.Context
	repository *repositoryMock
	controller todo.Controller
}

func (suite *ControllerTestSuite) SetupTest() {
	repository := new(repositoryMock)
	repository.On("Save", mock.Anything, mock.Anything).Return(nil)
	repository.On("List", mock.Anything, mock.Anything).Return([]todo.Todo{{}}, nil)
	repository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	suite.repository = repository
	suite.controller = todo.NewController(slog.New(slog.NewTextHandler(os.Stdout, nil)), repository)
	suite.ctx = context.Background()
}

func (suite *ControllerTestSuite) TestController_CreateNewTodo() {
	err := suite.controller.CreateNewTodo(suite.ctx, todo.Todo{
		Text: "clean kitchen",
		Done: false,
	})

	suite.NoError(err)
}

func (suite *ControllerTestSuite) TestController_CreateNewTodo_NoText() {
	err := suite.controller.CreateNewTodo(suite.ctx, todo.Todo{
		Text: "",
		Done: false,
	})

	suite.ErrorIs(err, todo.ErrNoText)
}

func (suite *ControllerTestSuite) TestController_CreateNewTodo_TextTooLong() {
	err := suite.controller.CreateNewTodo(suite.ctx, todo.Todo{
		Text: generateTextBySize(todo.MaxText + 1),
		Done: false,
	})

	suite.ErrorIs(err, todo.ErrTextTooLong)
}

func (suite *ControllerTestSuite) TestController_ListTodos() {
	todos, err := suite.controller.ListTodos(suite.ctx)

	suite.Len(todos, 1)
	suite.NoError(err)
}

func (suite *ControllerTestSuite) TestController_UpdateTodo() {
	oldTodo := todo.Todo{
		Text: "old",
		Done: false,
	}
	newTodo := todo.Todo{
		Text: "new",
		Done: false,
	}

	err := suite.controller.UpdateTodo(suite.ctx, oldTodo, newTodo)

	suite.NoError(err)
}

func (suite *ControllerTestSuite) TestController_UpdateTodo_sameTodo() {
	t := todo.Todo{
		Text: "todo",
		Done: false,
	}

	err := suite.controller.UpdateTodo(suite.ctx, t, t)

	suite.NoError(err)
}

func (suite *ControllerTestSuite) TestController_UpdateTodo_NoText() {
	oldTodo := todo.Todo{
		Text: "old",
		Done: false,
	}

	err := suite.controller.UpdateTodo(suite.ctx, oldTodo, todo.Todo{})

	suite.ErrorIs(err, todo.ErrNoText)
}

func (suite *ControllerTestSuite) TestController_UpdateTodo_TextTooLong() {
	oldTodo := todo.Todo{
		Text: "old",
		Done: false,
	}
	newTodo := todo.Todo{
		Text: generateTextBySize(todo.MaxText + 1),
		Done: false,
	}

	err := suite.controller.UpdateTodo(suite.ctx, oldTodo, newTodo)

	suite.ErrorIs(err, todo.ErrTextTooLong)
}

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}

func generateTextBySize(size int) string {
	var text string
	for i := 0; i < size; i++ {
		text += "a"
	}
	return text
}
