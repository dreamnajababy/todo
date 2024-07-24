package todo

type TodoRepository interface {
	GetTodoList() []Todo
	Save(todo Todo) Todo
}

type InMemoryTodoRepository struct {
	todos []Todo
}

func (r InMemoryTodoRepository) GetTodoList() []Todo {
	return r.todos
}

func (r InMemoryTodoRepository) Save(todo Todo) Todo {
	r.todos = append(r.todos, todo)
	return todo
}
