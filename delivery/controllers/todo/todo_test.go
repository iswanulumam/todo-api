package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-layered/entities"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"

	"todo-layered/delivery/controllers/auth"
	_authController "todo-layered/delivery/controllers/auth"
	_middleware "todo-layered/delivery/middleware"
	_authRepository "todo-layered/repository/auth"
)

func TestGet(t *testing.T) {
	// setting controller
	t.Run("TestGet", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/todo")

		// setup mocking data
		todoController := New(mockTodoRepository{})
		todoController.Get()(context)

		var response []entities.Todo
		json.Unmarshal([]byte(res.Body.String()), &response)
		assert.Equal(t, response[0].Title, "Eat Banana")
	})
}

func TestCreate(t *testing.T) {

	// var (
	// 	// for other operation
	// 	globalToken = ""
	// )

	t.Run("TestLogin", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"username": "admin",
			"password": "admin",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authRepository := _authRepository.New()
		authController := _authController.New(authRepository)

		authController.Login()(context)

		var response auth.LoginResponseFormat
		json.Unmarshal([]byte(res.Body.String()), &response)

		// globalToken = response.Token
		assert.Equal(t, 159, len(response.Token))

	})

	t.Run("TestCreate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"title": "eat lah",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		// localToken := globalToken
		localToken, _ := _middleware.CreateToken(1, "admin")
		fmt.Println("token", localToken)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", localToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo")

		// setup mocking data
		todoController := New(mockTodoRepository{})
		middleware.JWT([]byte("R4HASIA"))(todoController.Create())(context)

		var response entities.Todo
		json.Unmarshal([]byte(res.Body.String()), &response)
		fmt.Println("title", response.Title)
		assert.Equal(t, "Eat Blueberry", response.Title)
	})
}

// =========================== mocking ===========================

type mockTodoRepository struct{}

func (m mockTodoRepository) Get() ([]entities.Todo, error) {
	return []entities.Todo{
		{Title: "Eat Banana"},
		{Title: "Eat Apple"},
	}, nil
}

func (m mockTodoRepository) Create(entities.Todo) (entities.Todo, error) {
	return entities.Todo{Title: "Eat Blueberry"}, nil
}
