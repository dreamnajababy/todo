package domain

import (
	"errors"
	"sort"

	"github.com/google/uuid"
)

type InMemoryTodoRepository struct {
	todos []Todo
}

func (r InMemoryTodoRepository) GetTodoList(sortableColumns map[TodoSortedColumn]OrderBy) []Todo {
	switch {
	case sortableColumns[TITLE] == ASC:
		sort.Sort(ByTitleAsc(r.todos))
	case sortableColumns[TITLE] == DESC:
		sort.Sort(ByTitleDesc(r.todos))
	case sortableColumns[CREATED_AT] == ASC:
		sort.Sort(ByCreatedAtAsc(r.todos))
	case sortableColumns[CREATED_AT] == DESC:
		sort.Sort(ByCreatedAtDesc(r.todos))
	case sortableColumns[STATUS] == ASC:
		sort.Sort(ByStatusAsc(r.todos))
	case sortableColumns[STATUS] == DESC:
		sort.Sort(ByStatusDesc(r.todos))
	}
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
