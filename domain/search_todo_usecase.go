package domain

import "strings"

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
	todos := usecase.repository.GetTodoList(map[TodoSortedColumn]OrderBy{})
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
