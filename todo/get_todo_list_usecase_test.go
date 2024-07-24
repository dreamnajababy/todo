package todo

import (
	"reflect"
	"testing"
)

func TestGetTodoListUseCase(t *testing.T) {
	t.Run("return todo list", func(t *testing.T) {
		// arrange
		repository := InMemoryTodoRepository{
			todos: []Todo{
				CreateTestTodoData(),
				CreateTestTodoData(),
			},
		}
		want := repository.GetTodoList()
		useCase := NewGetTodoListUseCase(&repository)

		// act
		got := useCase.Execute()

		// assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
