package router

import (
	"todo-layered/delivery/controllers/todo"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, todoController *todo.TodoController) {
	// ------------------------------------------------------------------
	// Todo
	// ------------------------------------------------------------------
	e.GET("/todo", todoController.Get)
	e.POST("/todo", todoController.Create)
}
