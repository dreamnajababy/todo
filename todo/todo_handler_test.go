package todo

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

func createTestTodoData() *Todo {
	todo, _ := NewTodo(
		uuid.New().String(),
		"__TEST_TITLE__",
		"__TEST_DESCRIPTION__",
		time.Now().UTC().Format(time.RFC3339),
		base64.StdEncoding.EncodeToString([]byte("__TEST_IMAGE_URL__")),
		IN_PROGRESS,
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
			todos: []Todo{*createTestTodoData(), *createTestTodoData()},
		}
		want := repository.GetTodoList()
		act := PrepareTodoHandler(&repository)

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
