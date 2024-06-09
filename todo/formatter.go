package todo

import "time"

type TodoFormatter struct {
	ID       int       `json:"id"`
	Todo     string    `json:"todo"`
	DateTime time.Time `json:"date_time"`
	UserID   int       `json:"user_id"`
}


func FormatTodo(todo Todo) TodoFormatter {
<<<<<<< HEAD
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
=======
	
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
	
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	if len(todos) == 0 {
		return []TodoFormatter{}
	}
	
<<<<<<< HEAD
	// TODO 2
	var todosFormatter []TodoFormatter

	// TODO 3
=======
	var todosFormatter []TodoFormatter
	
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	for _, todo := range todos {
		var todoFormatter = FormatTodo(todo)
		todosFormatter = append(todosFormatter, todoFormatter)
	}

	return todosFormatter
}
