package todo

import (
	"encoding/json"
	"net/http"
)

// TODO: Change dependency to be usecase instead of repository and group according to the module handler
func PrepareTodoHandler(todoRepository TodoRepository) func(w http.ResponseWriter, r *http.Request) {
	getTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		todos := todoRepository.GetTodoList()
		json.NewEncoder(w).Encode(todos)
		w.WriteHeader(http.StatusOK)
	}
	return getTodoHandler
}
