package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TodoRepository interface {
	GetTodoList() []Todo
	GetById(id uuid.UUID) (Todo, error)
	Save(todo Todo) Todo
}

type InMemoryTodoRepository struct {
	todos []Todo
}

func (r InMemoryTodoRepository) GetTodoList() []Todo {
	return r.todos
}

func (r *InMemoryTodoRepository) Save(todo Todo) Todo {
	r.todos = append(r.todos, todo)
	return todo
}

func (r InMemoryTodoRepository) GetById(id uuid.UUID) (Todo, error) {
	for _, todo := range r.todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return Todo{}, errors.New("todo not found")
}
