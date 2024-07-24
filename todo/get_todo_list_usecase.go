package todo

type GetTodoListUseCase struct {
	todoRepository TodoRepository
}

func (usecase GetTodoListUseCase) Execute() []Todo {
	return usecase.todoRepository.GetTodoList()
}

func NewGetTodoListUseCase(todoRepository TodoRepository) GetTodoListUseCase {
	return GetTodoListUseCase{todoRepository: todoRepository}
}
