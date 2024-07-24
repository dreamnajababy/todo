package todo

import "time"

type CreateTodoRequest struct {
	Id          string
	Title       string
	Description string
	CreatedAt   string
	Image       string
	Status      Status
}
type CreateTodoUseCase struct {
	repository TodoRepository
}

func NewCreateTodoUseCase(repository TodoRepository) CreateTodoUseCase {
	return CreateTodoUseCase{repository: repository}
}

func (usecase CreateTodoUseCase) Execute(request CreateTodoRequest) (Todo, error) {
	parsedTime, _ := time.Parse(time.RFC3339, request.CreatedAt)
	todo, err := NewTodo(
		request.Id,
		request.Title,
		request.Description,
		parsedTime.Format(time.RFC3339),
		request.Image,
		request.Status,
	)
	if err != nil {
		return Todo{}, err
	}
	createdTodo := usecase.repository.Save(todo)
	return createdTodo, nil
}
