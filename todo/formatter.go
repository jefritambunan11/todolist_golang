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
	
	var todoFormatter = TodoFormatter{}
	
	todoFormatter.ID = todo.ID
	todoFormatter.Todo = todo.Todo
	todoFormatter.DateTime = todo.DateTime
	todoFormatter.UserID = todo.UserID	
	
	return todoFormatter
}

func FormatTodoDetail(todo Todo) TodoFormatter {
	
	var todoFormatter = TodoFormatter{}	

	todoFormatter.ID = todo.ID
	todoFormatter.UserID = todo.UserID
	todoFormatter.Todo = todo.Todo
	todoFormatter.DateTime = todo.DateTime
	
	var todoUserFormatter = TodoUserFormatter{}
	todoUserFormatter.ID = todo.User.ID
	
	todoFormatter.User = todoUserFormatter
	
	return todoFormatter	
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
