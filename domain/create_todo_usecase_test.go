package domain

import (
	"encoding/base64"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateTodoUseCase(t *testing.T) {
	t.Run("Create new todo", func(t *testing.T) {
		// arrange
		id := uuid.New()
		repository := InMemoryTodoRepository{
			todos: []Todo{},
		}
		request := CreateTodoRequest{
			Id:          id.String(),
			Title:       "__TEST_TITLE__",
			Description: "__TEST_DESCRIPTION__",
			CreatedAt:   "2018-02-16T05:15:37Z",
			Image:       base64.StdEncoding.EncodeToString([]byte("__TEST_IMAGE__")),
			Status:      IN_PROGRESS,
		}
		parsedTime, _ := time.Parse(time.RFC3339, request.CreatedAt)
		want := Todo{
			Id:          id,
			Title:       request.Title,
			Description: request.Description,
			CreatedAt:   parsedTime,
			Image:       request.Image,
			Status:      request.Status,
		}

		useCase := NewCreateTodoUseCase(&repository)

		got, err := useCase.Execute(request)

		if err != nil {
			t.Errorf("Error while creating todo: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Create new todo with something wrong and should return some error", func(t *testing.T) {
		// arrange
		repository := InMemoryTodoRepository{
			todos: []Todo{},
		}
		request := CreateTodoRequest{
			Id:          "__INVALID_ID__",
			Title:       "__TEST_TITLE__",
			Description: "__TEST_DESCRIPTION__",
			CreatedAt:   "2018-02-16T05:15:37Z",
			Image:       base64.StdEncoding.EncodeToString([]byte("__TEST_IMAGE__")),
			Status:      IN_PROGRESS,
		}
		want := "id must be UUID format"

		useCase := NewCreateTodoUseCase(&repository)

		_, err := useCase.Execute(request)

		if err == nil {
			t.Error("Error should be returned when creating todo with invalid id")
		}
		if err.Error() != want {
			t.Errorf("got %v want %v", err.Error(), want)
		}
	})
}
