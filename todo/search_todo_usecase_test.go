package todo

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

type SearchTodoUseCase struct {
	repository TodoRepository
}

type TodoSearchableColumn string

type SearchTodoRequest struct {
	keyword string
	columns []TodoSearchableColumn
}

func NewSearchTodoUseCase(repository TodoRepository) SearchTodoUseCase {
	return SearchTodoUseCase{repository: repository}
}

func (usecase SearchTodoUseCase) searchByTitle(todos []Todo, result []Todo, keyword string) []Todo {
	for _, todo := range todos {
		if strings.Contains(todo.Title, keyword) {
			result = append(result, todo)
		}
	}
	return result
}

func (usecase SearchTodoUseCase) searchByDescription(todos []Todo, result []Todo, keyword string) []Todo {
	for _, todo := range todos {
		if strings.Contains(todo.Description, keyword) {
			result = append(result, todo)
		}
	}
	return result
}

func (usecase SearchTodoUseCase) Execute(request SearchTodoRequest) []Todo {
	todos := usecase.repository.GetTodoList()
	var result []Todo
	for _, column := range request.columns {
		switch column {
		case "Title":
			result = usecase.searchByTitle(todos, result, request.keyword)
		case "Description":
			result = usecase.searchByDescription(todos, result, request.keyword)
		}
	}
	return result
}

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
