package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTodoListUseCase(t *testing.T) {
	t.Run("Return sorted todo list by title with ascending", func(t *testing.T) {
		// arrange
		firstTodo := CreateTestTodoData(
			"__TEST_TITLE_1__",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE_2__",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				secondTodo,
				firstTodo,
			},
		}
		request := GetTodoListRequest{
			sortableColumns: map[TodoSortedColumn]OrderBy{
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
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE_2__",
			time.Now().UTC().Format(time.RFC3339),
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		request := GetTodoListRequest{
			sortableColumns: map[TodoSortedColumn]OrderBy{
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
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				secondTodo,
				firstTodo,
			},
		}
		request := GetTodoListRequest{
			sortableColumns: map[TodoSortedColumn]OrderBy{
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
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		request := GetTodoListRequest{
			sortableColumns: map[TodoSortedColumn]OrderBy{
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
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			"__TEST_DESCRIPTION__",
			COMPLETE,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				firstTodo,
				secondTodo,
			},
		}
		request := GetTodoListRequest{
			sortableColumns: map[TodoSortedColumn]OrderBy{
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
			"__TEST_DESCRIPTION__",
			IN_PROGRESS,
		)
		secondTodo := CreateTestTodoData(
			"__TEST_TITLE__",
			"2018-02-16T05:15:37Z",
			"__TEST_DESCRIPTION__",
			COMPLETE,
		)
		repository := InMemoryTodoRepository{
			todos: []Todo{
				secondTodo,
				firstTodo,
			},
		}
		request := GetTodoListRequest{
			sortableColumns: map[TodoSortedColumn]OrderBy{
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
					"__TEST_DESCRIPTION__",
					IN_PROGRESS,
				),
				CreateTestTodoData(
					"__TEST_TITLE__",
					time.Now().UTC().Format(time.RFC3339),
					"__TEST_DESCRIPTION__",
					IN_PROGRESS,
				),
			},
		}
		request := GetTodoListRequest{}
		want := repository.GetTodoList(map[TodoSortedColumn]OrderBy{})
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute(request)

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
