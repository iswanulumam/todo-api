package todo

import (
	"gorm.io/gorm"

	"todo-layered/entities"
)

type TodoRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (tr *TodoRepository) Get() ([]entities.Todo, error) {
	var todo []entities.Todo
	if err := tr.db.Find(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (tr *TodoRepository) Create(todo entities.Todo) (entities.Todo, error) {
	if err := tr.db.Save(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}
