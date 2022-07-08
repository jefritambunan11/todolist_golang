package todo

import "time"

type TodoFormatter struct {
	ID       int       `json:"id"`
	Todo     string    `json:"todo"`
	DateTime time.Time `json:"date_time"`
	UserID   int       `json:"user_id"`

	User TodoUserFormatter `json:"user"`
}

type TodoUserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatTodo(todo Todo) TodoFormatter {
	todoFormatter := TodoFormatter{}
	todoFormatter.ID = todo.ID
	todoFormatter.Todo = todo.Todo
	todoFormatter.DateTime = todo.DateTime
	todoFormatter.UserID = todo.UserID

	return todoFormatter
}

func FormatTodoDetail(todo Todo) TodoFormatter {
	t := TodoFormatter{}
	t.ID = todo.ID
	t.UserID = todo.UserID
	t.Todo = todo.Todo
	t.DateTime = todo.DateTime

	todoUserFormatter := TodoUserFormatter{}
	todoUserFormatter.ID = todo.User.ID

	t.User = todoUserFormatter

	return t
}

func FormatTodos(todos []Todo) []TodoFormatter {
	if len(todos) == 0 {
		return []TodoFormatter{}
	}

	var todosFormatter []TodoFormatter

	for _, todo := range todos {
		todoFormatter := FormatTodo(todo)
		todosFormatter = append(todosFormatter, todoFormatter)
	}

	return todosFormatter
}
