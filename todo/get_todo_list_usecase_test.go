package todo

import (
	"reflect"
	"testing"
	"time"
)

// Can search the data by Title or Description fields
// The TODO application can UPDATE a task with the following requirements
// Can update a task by ID field
// Can update Title, Description, Date, Image, and Status fields corresponding to the requirements from the CREATE feature

func TestGetTodoListUseCase(t *testing.T) {
	t.Run("Return sorted todo list by title with ascending", func(t *testing.T) {
		// arrange
		firstTodo := CreateTestTodoData(
			"__TEST_TITLE_1__",
			time.Now().UTC().Format(time.RFC3339),
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE_2__",
			time.Now().UTC().Format(time.RFC3339),
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				secondTodo,
				firstTodo,
			},
		}
		request := GetTodoListRequest{
			columns: map[TodoSortedColumn]OrderBy{
				TITLE: ASC,
			},
		}
		want := []Todo{firstTodo, secondTodo}
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Return sorted todo list by title with descending", func(t *testing.T) {
		// arrange
		firstTodo := CreateTestTodoData(
			"__TEST_TITLE_1__",
			time.Now().UTC().Format(time.RFC3339),
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE_2__",
			time.Now().UTC().Format(time.RFC3339),
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		request := GetTodoListRequest{
			columns: map[TodoSortedColumn]OrderBy{
				TITLE: DESC,
			},
		}
		want := []Todo{secondTodo, firstTodo}
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Return sorted todo list by CreatedAt with ascending", func(t *testing.T) {
		// arrange
		firstTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-01-16T05:15:37Z",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				secondTodo,
				firstTodo,
			},
		}
		request := GetTodoListRequest{
			columns: map[TodoSortedColumn]OrderBy{
				CREATED_AT: ASC,
			},
		}
		want := []Todo{firstTodo, secondTodo}
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Return sorted todo list by CreatedAt with descending", func(t *testing.T) {
		// arrange
		firstTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-01-16T05:15:37Z",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		request := GetTodoListRequest{
			columns: map[TodoSortedColumn]OrderBy{
				CREATED_AT: DESC,
			},
		}
		want := []Todo{secondTodo, firstTodo}
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Return sorted todo list by Status with ascending", func(t *testing.T) {
		// arrange
		firstTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			COMPLETE,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		request := GetTodoListRequest{
			columns: map[TodoSortedColumn]OrderBy{
				STATUS: ASC,
			},
		}
		want := []Todo{secondTodo, firstTodo}
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Return sorted todo list by Status with descending", func(t *testing.T) {
		// arrange
		firstTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			COMPLETE,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				secondTodo,
				firstTodo,
			},
		}
		request := GetTodoListRequest{
			columns: map[TodoSortedColumn]OrderBy{
				STATUS: DESC,
			},
		}
		want := []Todo{firstTodo, secondTodo}
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("return todo list", func(t *testing.T) {
		// arrange
		repository := InMemoryTodoRepository{
			todos: []Todo{
				CreateTestTodoData(
					"__TEST_TITLE__",
					time.Now().UTC().Format(time.RFC3339),
					IN_PROGRESS,
				),
				CreateTestTodoData(
					"__TEST_TITLE__",
					time.Now().UTC().Format(time.RFC3339),
					IN_PROGRESS,
				),
			},
		}
		request := GetTodoListRequest{}
		want := repository.GetTodoList()
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
