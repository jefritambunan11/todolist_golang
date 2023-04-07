package todo

import "time"

type TodoFormatter struct {
	ID       int       `json:"id"`
	Todo     string    `json:"todo"`
	DateTime time.Time `json:"date_time"`
	UserID   int       `json:"user_id"`
}


func FormatTodo(todo Todo) TodoFormatter {
	
	var todoFormatter = TodoFormatter{}
	
	todoFormatter.ID = todo.ID
	todoFormatter.Todo = todo.Todo
	todoFormatter.DateTime = todo.DateTime
	todoFormatter.UserID = todo.UserID	
	
	return todoFormatter
}


func FormatTodoDetail(todo Todo) TodoFormatter  {
	
	var todoFormatter = TodoFormatter{}	

	todoFormatter.ID = todo.ID
	todoFormatter.UserID = todo.UserID
	todoFormatter.Todo = todo.Todo
	todoFormatter.DateTime = todo.DateTime
	
	return todoFormatter	
}


func FormatTodos(todos []Todo) []TodoFormatter  {
	
	if len(todos) == 0 {
		return []TodoFormatter{}
	}
	
	var todosFormatter []TodoFormatter
	
	for _, todo := range todos {
		var todoFormatter = FormatTodo(todo)
		todosFormatter = append(todosFormatter, todoFormatter)
	}

	return todosFormatter
}
