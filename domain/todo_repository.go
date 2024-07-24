package domain

import "github.com/google/uuid"

type TodoRepository interface {
	GetTodoList(sortableColumns map[TodoSortedColumn]OrderBy) []Todo
	GetById(id uuid.UUID) (Todo, error)
	Save(todo Todo) Todo
}
