package domain

import (
	"encoding/json"
	"net/http"
)

// TODO: Change dependency to be usecase instead of repository and group according to the module handler
func PrepareTodoHandler(getTodoListUseCase GetTodoListUseCase) func(w http.ResponseWriter, r *http.Request) {
	getTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		request := GetTodoListRequest{}
		todos := getTodoListUseCase.Execute(request)
		json.NewEncoder(w).Encode(todos)
		w.WriteHeader(http.StatusOK)
	}
	return getTodoHandler
}
