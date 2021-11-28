package auth

import (
	"net/http"
	"todo-layered/delivery/common"
	authRepo "todo-layered/repository/auth"

	echo "github.com/labstack/echo/v4"
)

type AuthController struct {
	repository authRepo.Auth
}

func New(auth authRepo.Auth) *AuthController {
	return &AuthController{
		repository: auth,
	}
}

func (a AuthController) Login(c echo.Context) error {
	var loginRequest LoginRequestFormat

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequest())
	}

	token, err := a.repository.Login(loginRequest.Username, loginRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.BadRequest())
	}

	loginResponse := LoginResponseFormat{
		Token: token,
	}

	return c.JSON(http.StatusOK, loginResponse)

}
