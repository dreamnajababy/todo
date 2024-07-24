package todo

import (
	"reflect"
	"testing"
)

// Can sort the data by Title or Date or Status fields
// Can search the data by Title or Description fields
// The TODO application can UPDATE a task with the following requirements
// Can update a task by ID field
// Can update Title, Description, Date, Image, and Status fields corresponding to the requirements from the CREATE feature

func TestGetTodoListUseCase(t *testing.T) {
	t.Run("Return sorted todo list by title with ascending", func(t *testing.T) {
		// arrange
		repository := InMemoryTodoRepository{
			todos: []Todo{
				CreateTestTodoData(),
				CreateTestTodoData(),
			},
		}
		request := GetTodoListRequest{
			columns: map[TodoSortedColumn]OrderBy{
				TITLE: ASC,
			},
		}
		want := repository.GetTodoList()
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
				CreateTestTodoData(),
				CreateTestTodoData(),
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
