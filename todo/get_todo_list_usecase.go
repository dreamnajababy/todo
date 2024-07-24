package todo

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

func (usecase GetTodoListUseCase) Execute(request GetTodoListRequest) []Todo {
	return usecase.todoRepository.GetTodoList()
}

func NewGetTodoListUseCase(todoRepository TodoRepository) GetTodoListUseCase {
	return GetTodoListUseCase{todoRepository: todoRepository}
}
