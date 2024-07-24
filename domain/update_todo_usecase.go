package domain

import (
	"errors"

	"github.com/google/uuid"
)

type UpdateTodoUseCase struct {
	repository TodoRepository
}

func NewUpdateTodoUsecase(repository TodoRepository) UpdateTodoUseCase {
	return UpdateTodoUseCase{repository: repository}
}

func (usecase UpdateTodoUseCase) Execute(request UpdateTodoRequest) (Todo, error) {
	_, err := usecase.repository.GetById(uuid.MustParse(request.Id))
	if err != nil {
		return Todo{}, errors.New("todo not found")
	}
	todo, err := NewTodo(
		request.Id,
		request.Title,
		request.Description,
		request.CreatedAt,
		request.Image,
		request.Status,
	)
	if err != nil {
		return Todo{}, err
	}
	updatedTodo := usecase.repository.Save(todo)
	return updatedTodo, nil
}

type UpdateTodoRequest struct {
	Id          string
	Title       string
	Description string
	CreatedAt   string
	Image       string
	Status      Status
}
