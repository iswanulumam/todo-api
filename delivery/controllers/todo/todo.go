package todo

import (
	"fmt"
	"net/http"
	"todo-layered/delivery/common"
	"todo-layered/entities"
	todoRepo "todo-layered/repository/todo"

	"github.com/golang-jwt/jwt"
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

func (tc TodoController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo, err := tc.repository.Get()

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc TodoController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)

		if !user.Valid {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}

		claims := user.Claims.(jwt.MapClaims)

		//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
		//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
		userID, ok := claims["id"].(float64)

		fmt.Println("inject jwt with testing", int(userID))

		if !ok {
			return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
		}

		var todoRequest TodoRequestFormat

		if err := c.Bind(&todoRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		todo := entities.Todo{
			Title: todoRequest.Title,
		}

		response, err := tc.repository.Create(todo)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, response)
	}
}
