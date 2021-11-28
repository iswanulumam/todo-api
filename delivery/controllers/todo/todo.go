package todo

import (
	"net/http"
	"todo-layered/delivery/common"
	"todo-layered/entities"
	todoRepo "todo-layered/repository/todo"

	echo "github.com/labstack/echo/v4"
)

type TodoController struct {
	repository todoRepo.Todo
}

func New(todo todoRepo.Todo) *TodoController {
	return &TodoController{
		repository: todo,
	}
}

func (tc TodoController) Get(c echo.Context) error {
	todo, err := tc.repository.Get()

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequest())
	}

	return c.JSON(http.StatusOK, todo)
}

func (tc TodoController) Create(c echo.Context) error {
	var todoRequest TodoRequestFormat

	if err := c.Bind(&todoRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequest())
	}

	todo := entities.Todo{
		Title: todoRequest.Title,
	}

	_, err := tc.repository.Create(todo)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerError())
	}

	return c.JSON(http.StatusOK, common.SuccessOperation())
}
