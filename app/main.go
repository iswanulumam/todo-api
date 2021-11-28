package main

import (
	"todo-layered/config"

	_authRepo "todo-layered/repository/auth"
	_todoRepo "todo-layered/repository/todo"

	_authController "todo-layered/delivery/controllers/auth"
	_todoController "todo-layered/delivery/controllers/todo"

	"todo-layered/util"

	"todo-layered/delivery/router"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDriver(config)

	//initiate user model
	todoRepo := _todoRepo.New(db)
	authRepo := _authRepo.New()

	//initiate user controller
	todoController := _todoController.New(todoRepo)
	authController := _authController.New(authRepo)

	//create echo http
	e := echo.New()

	//register API path and controller
	router.RegisterPath(e, todoController, authController)

	// run server
	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
