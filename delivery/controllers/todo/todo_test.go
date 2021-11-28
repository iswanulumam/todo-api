package todo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-layered/entities"

	echo "github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
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
		todoController.Get(context)

		var result []entities.Todo
		json.Unmarshal([]byte(res.Body.String()), &result)

		assert.Equal(t, result[0].Title, "Eat Banana")
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
	return entities.Todo{
		Title: "Todo",
	}, nil
}
