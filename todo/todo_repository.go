package todo

type TodoRepository interface {
	GetTodoList() []Todo
}

type InMemoryTodoRepository struct {
	todos []Todo
}

func (r *InMemoryTodoRepository) GetTodoList() []Todo {
	return r.todos
}
