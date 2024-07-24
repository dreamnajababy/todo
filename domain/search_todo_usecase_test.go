package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestSearchTodoUseCase(t *testing.T) {
	// Can search the data by Title or Description fields
	t.Run("Return todo list with title contains 'test'", func(t *testing.T) {
		firstTodo := CreateTestTodoData(
			"lorem ipsum test",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"lorem upsum",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		want := []Todo{
			firstTodo,
		}
		request := SearchTodoRequest{
			keyword: "test",
			columns: []TodoSearchableColumn{"Title"},
		}

		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		useCase := NewSearchTodoUseCase(&repository)
		got := useCase.Execute(request)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Return todo list with description contains 'test'", func(t *testing.T) {
		firstTodo := CreateTestTodoData(
			"lorem ipsum test",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"lorem upsum",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		want := []Todo{
			firstTodo,
		}
		request := SearchTodoRequest{
			keyword: "test",
			columns: []TodoSearchableColumn{"Title"},
		}

		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		useCase := NewSearchTodoUseCase(&repository)
		got := useCase.Execute(request)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
