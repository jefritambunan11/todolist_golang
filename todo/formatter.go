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
	// TODO 
	// 1. Pake struct TodoFormatter ke todoFormatter
	// 2. set data field di struct dari struct todo
	// 3. return struct TodoFormatter
	
	// TODO 1
	var todoFormatter = TodoFormatter{}

	// TODO 2
	todoFormatter.ID = todo.ID
	todoFormatter.Todo = todo.Todo
	todoFormatter.DateTime = todo.DateTime
	todoFormatter.UserID = todo.UserID
	
	// TODO 3
	return todoFormatter
}

func FormatTodoDetail(todo Todo) TodoFormatter {
	// TODO
	// 1. pake struct TodoFormatter
	// 2. isi data struct TodoFormatter
	// 3. pake struct todoUserFormatter dan isi field ID dari todo.User.ID
	// 4. isikan field User dari TodoFormatter pake todoUserFormatter
	// 5. return struct TodoFormatter
	
	// TODO 1
	var todoFormatter = TodoFormatter{}
	
	// TODO 2
	todoFormatter.ID = todo.ID
	todoFormatter.UserID = todo.UserID
	todoFormatter.Todo = todo.Todo
	todoFormatter.DateTime = todo.DateTime

	// TODO 3
	var todoUserFormatter = TodoUserFormatter{}
	todoUserFormatter.ID = todo.User.ID

	// TODO 4
	todoFormatter.User = todoUserFormatter

	// TODO 5
	return todoFormatter
}

func FormatTodos(todos []Todo) []TodoFormatter {
	// TODO 
	// 1. jika multi struct kosong return multi struct TodoFormatter
	// 2. pake multi struct TodoFormatter di todosFormatter
	// 3. 	
	
	// TODO 1
	if len(todos) == 0 {
		return []TodoFormatter{}
	}
	
	// TODO 2
	var todosFormatter []TodoFormatter

	// TODO 3
	for _, todo := range todos {
		todoFormatter := FormatTodo(todo)
		todosFormatter = append(todosFormatter, todoFormatter)
	}

	return todosFormatter
}
