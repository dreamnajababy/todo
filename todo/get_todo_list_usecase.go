package todo

import (
	"sort"
)

type GetTodoListUseCase struct {
	todoRepository TodoRepository
}

const (
	TITLE  TodoSortedColumn = "Title"
	DATE   TodoSortedColumn = "Date"
	STATUS TodoSortedColumn = "Status"
)
const (
	ASC  OrderBy = "ASC"
	DESC OrderBy = "DESC"
)

type TodoSortedColumn string
type OrderBy string

type GetTodoListRequest struct {
	columns map[TodoSortedColumn]OrderBy
}

func thereIsNo(columns map[TodoSortedColumn]OrderBy) bool {
	return len(columns) == 0
}

func (usecase GetTodoListUseCase) Execute(request GetTodoListRequest) []Todo {
	if thereIsNo(request.columns) {
		return usecase.todoRepository.GetTodoList()
	}
	todos := usecase.todoRepository.GetTodoList()
	switch {
	case request.columns[TITLE] == ASC:
		sort.Sort(ByTitleAsc(todos))
	case request.columns[TITLE] == DESC:
		sort.Sort(ByTitleDesc(todos))
	case request.columns[DATE] == ASC:
		sort.Sort(ByCreatedAtAsc(todos))
	case request.columns[DATE] == DESC:
		sort.Sort(ByCreatedAtDesc(todos))
	case request.columns[STATUS] == ASC:
		sort.Sort(ByStatusAsc(todos))
	}
	return todos
}

func NewGetTodoListUseCase(todoRepository TodoRepository) GetTodoListUseCase {
	return GetTodoListUseCase{todoRepository: todoRepository}
}
