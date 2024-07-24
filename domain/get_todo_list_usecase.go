package domain

type GetTodoListUseCase struct {
	todoRepository TodoRepository
}

const (
	TITLE      TodoSortedColumn = "Title"
	CREATED_AT TodoSortedColumn = "Date"
	STATUS     TodoSortedColumn = "Status"
)
const (
	ASC  OrderBy = "ASC"
	DESC OrderBy = "DESC"
)

type TodoSortedColumn string
type OrderBy string

type GetTodoListRequest struct {
	sortableColumns map[TodoSortedColumn]OrderBy
}

func thereIsNo(columns map[TodoSortedColumn]OrderBy) bool {
	return len(columns) == 0
}

func (usecase GetTodoListUseCase) Execute(request GetTodoListRequest) []Todo {
	if thereIsNo(request.sortableColumns) {
		return usecase.todoRepository.GetTodoList(request.sortableColumns)
	}
	return usecase.todoRepository.GetTodoList(request.sortableColumns)
}

func NewGetTodoListUseCase(todoRepository TodoRepository) GetTodoListUseCase {
	return GetTodoListUseCase{todoRepository: todoRepository}
}
