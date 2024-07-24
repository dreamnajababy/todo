package todo

import (
	"encoding/base64"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUpdateTodoUsecase(t *testing.T) {
	t.Run("should update todo succesfully", func(t *testing.T) {
		id := uuid.New()
		createdAt, err := time.Parse(time.RFC3339, "2018-02-16T05:15:37Z")
		if err != nil {
			t.Errorf("Error parsing time: %v", err)
		}
		request := UpdateTodoRequest{
			Id:          id.String(),
			Title:       "__TEST_TITLE_UPDATED__",
			Description: "__TEST_DESCRIPTION_UPDATED__",
			CreatedAt:   "2018-02-17T05:15:37Z",
			Image:       base64.StdEncoding.EncodeToString([]byte("__TEST_IMAGE_UPDATED__")),
			Status:      COMPLETE,
		}
		want, err := NewTodo(request.Id, request.Title, request.Description, request.CreatedAt, request.Image, request.Status)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				{
					Id:          id,
					Title:       "__TEST_TITLE__",
					Description: "__TEST_DESCRIPTION__",
					CreatedAt:   createdAt,
					Image:       base64.StdEncoding.EncodeToString([]byte("__TEST_IMAGE__")),
					Status:      IN_PROGRESS,
				},
			},
		}
		usecase := NewUpdateTodoUsecase(&repository)
		if err != nil {
			t.Errorf("Error creating expected todo: %v", err)
		}

		got, err := usecase.Execute(request)

		if err != nil {
			t.Errorf("Error updating todo: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected todo to be %v but got %v", want, got)
		}
	})
}
