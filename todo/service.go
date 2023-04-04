package todo

import (
	"errors"
	
)

type Service interface {
	GetTodos(userID int, _page_number_ int) ([]Todo, error)
	GetTodoByID(input GetTodoDetailInput) (Todo, error)
	CreateTodo(input CreateTodoInput) (Todo, error)
	UpdateTodo(inputID GetTodoDetailInput, inputData CreateTodoInput) (Todo, error)
	DeleteTodo(inputID GetTodoDetailInput, inputData CreateTodoInput) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTodos(userID int, _page_number_ int) ([]Todo, error) {
	
	if userID != 0 {
		var todos, err = s.repository.FindByUserID(userID)
		
		if err != nil {
			return todos, err
		}
		
		return todos, nil
	}
	
	
	var todos, err = s.repository.FindAll(_page_number_)
	if err != nil {
		return todos, err
	}
	
	
	return todos, nil
}

func (s *service) GetTodoByID(input GetTodoDetailInput) (Todo, error) {
	
	var todo, err = s.repository.FindByID(input.ID)
	if err != nil {
		return todo, err
	}
	
	return todo, nil
}

func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	
	var todo = Todo{}
	todo.Todo = input.Todo
	todo.DateTime = input.DateTime
	todo.UserID = input.User.ID

	
	var newTodo, err = s.repository.Save(todo)
	if err != nil {
		return newTodo, err
	}
	
	
	return newTodo, nil

}

func (s *service) UpdateTodo(inputID GetTodoDetailInput, inputData CreateTodoInput) (Todo, error) {
	
	var todo, err = s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}
	
	
	if todo.UserID != inputData.User.ID {
		return todo, errors.New("todo nya bukan milik si user")
	}
	
	
	todo.UserID = inputData.User.ID
	todo.Todo = inputData.Todo
	todo.DateTime = inputData.DateTime
	
	
	updateTodo, err := s.repository.Update(todo)
	if err != nil {
		return updateTodo, err
	}
	
	
	return updateTodo, nil

}

func (s *service) DeleteTodo(inputID GetTodoDetailInput, inputData CreateTodoInput) (Todo, error) {
	
	var todo, err = s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}
	
	
	if todo.UserID != inputData.User.ID {
		return todo, errors.New("todo nya bukan milik si user")
	}
	
	
	var deleteTodo, err2 = s.repository.Delete(todo)
	if err2 != nil {
		return deleteTodo, err
	}
	
	
	return deleteTodo, nil

}
