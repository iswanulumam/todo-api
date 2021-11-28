package todo

import (
	"testing"
	"todo-layered/config"
	"todo-layered/entities"
	"todo-layered/util"

	"github.com/magiconair/properties/assert"
)

var (
	todoTitle = "Eat banana"
)

func TestMain(t *testing.T) {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDriver(config)

	// cleaning data before testing
	db.Migrator().DropTable(&entities.Todo{})
	db.AutoMigrate(&entities.Todo{})

	todoRepository := New(db)

	TodoCreate(t, todoRepository)
	TodoGet(t, todoRepository)
}

func TodoCreate(t *testing.T, repository *TodoRepository) {
	t.Run("TodoCreate", func(t *testing.T) {
		todo := entities.Todo{
			Title: todoTitle,
		}
		result, _ := repository.Create(todo)
		assert.Equal(t, result.Title, todo.Title)
	})
}

func TodoGet(t *testing.T, repository *TodoRepository) {
	t.Run("TodoGet", func(t *testing.T) {
		result, _ := repository.Get()
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0].Title, todoTitle)
	})
}
