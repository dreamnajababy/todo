package domain

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func CreateTestTodoData(title string, date string, description string, status Status) Todo {
	parsedDate, _ := time.Parse(time.RFC3339, date)
	todo, _ := NewTodo(
		uuid.New().String(),
		title,
		description,
		parsedDate.Format(time.RFC3339),
		base64.StdEncoding.EncodeToString([]byte("__TEST_IMAGE_URL__")),
		status,
	)
	return todo
}
func TestHTTPGetTodoList(t *testing.T) {
	t.Run("return todo list and get 200", func(t *testing.T) {
		// arrange
		var got []Todo
		request, _ := http.NewRequest(http.MethodGet, "/todos", nil)
		response := httptest.NewRecorder()
		repository := InMemoryTodoRepository{
			todos: []Todo{CreateTestTodoData(
				"__TEST_TITLE__",
				time.Now().UTC().Format(time.RFC3339),
				"__TEST_DESCRIPTION__",
				IN_PROGRESS,
			), CreateTestTodoData(
				"__TEST_TITLE__",
				time.Now().UTC().Format(time.RFC3339),
				"__TEST_DESCRIPTION__",
				IN_PROGRESS,
			)},
		}
		useCase := NewGetTodoListUseCase(&repository)
		want := repository.GetTodoList()
		act := PrepareTodoHandler(useCase)

		// act
		act(response, request)
		err := json.NewDecoder(response.Body).Decode(&got)

		// assert
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Todo, '%v'", response.Body, err)
		}
		if response.Code != http.StatusOK {
			t.Errorf("got %d, want %d", response.Code, http.StatusOK)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
