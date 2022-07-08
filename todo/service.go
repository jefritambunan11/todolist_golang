package todo

import (
	"errors"
	// "todolist/user"
	// "fmt"
	// "github.com/gosimple/slug"
)

type Service interface {
	GetTodos(userID int) ([]Todo, error)
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

func (s *service) GetTodos(userID int) ([]Todo, error) {
	if userID != 0 {
		todos, err := s.repository.FindByUserID(userID)
		if err != nil {
			return todos, err
		}
		return todos, nil
	}

	todos, err := s.repository.FindAll()
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *service) GetTodoByID(input GetTodoDetailInput) (Todo, error) {
	todo, err := s.repository.FindByID(input.ID)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	todo := Todo{}
	todo.Todo = input.Todo
	todo.DateTime = input.DateTime
	todo.UserID = input.User.ID

	newTodo, err := s.repository.Save(todo)
	if err != nil {
		return newTodo, err
	}

	return newTodo, nil

}

func (s *service) UpdateTodo(inputID GetTodoDetailInput, inputData CreateTodoInput) (Todo, error) {
	todo, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}

	if todo.UserID != inputData.User.ID {
		return todo, errors.New("todo nya bukan milik si user")
	}

	// user_id ambil dari middleware
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
	todo, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}

	if todo.UserID != inputData.User.ID {
		return todo, errors.New("todo nya bukan milik si user")
	}

	deleteTodo, err := s.repository.Delete(todo)

	if err != nil {
		return deleteTodo, err
	}

	return deleteTodo, nil

}
