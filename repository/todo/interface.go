package todo

import "todo-layered/entities"

type Todo interface {
	Get() ([]entities.Todo, error)
	Create(entities.Todo) (entities.Todo, error)
}
