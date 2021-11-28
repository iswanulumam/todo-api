package router

import (
	"todo-layered/delivery/controllers/auth"
	"todo-layered/delivery/controllers/todo"
	"todo-layered/delivery/middleware"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	todoController *todo.TodoController,
	authController *auth.AuthController) {

	// Login
	e.POST("/login", authController.Login)

	// Todo
	e.GET("/todo", todoController.Get)
	e.POST("/todo", todoController.Create, middleware.JWTMiddleware())
}
