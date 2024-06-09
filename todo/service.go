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
	// TODO
	// 1. filter jika user tidak 0 panggil method FindByUserID dan return nya data 1 baris data dari tabel uuser
	// 2. panggil method FindAll dan dapat semua data di tabel Todo
	// 3. Return seluruh tabel todo
	
	// TODO 1
	if userID != 0 {
		var todos, err = s.repository.FindByUserID(userID)
		
		if err != nil {
			return todos, err
		}
		
		return todos, nil
	}
	
	// TODO 2
	var todos, err = s.repository.FindAll()
	if err != nil {
		return todos, err
	}
	
	// TODO 3
	return todos, nil
}

func (s *service) GetTodoByID(input GetTodoDetailInput) (Todo, error) {
	// TODO	
	// 1. panggil method FindByID dan dapat data 1 baris dari tabel todo
	// 2. return 1 baris data tersebut
	
	// TODO 1
	var todo, err = s.repository.FindByID(input.ID)
	if err != nil {
		return todo, err
	}

	// TODO 2
	return todo, nil
}

func (s *service) CreateTodo(input CreateTodoInput) (Todo, error) {
	// TODO
	// 1. var todo pake struct Todo dan isi data field nya sesuai params input yg bertipe struct 
	// 2. panggil method Save dan oper struct todo nya
	// 3. return data struct yg baru di insert
	
	// TODO 1
	var todo = Todo{}
	todo.Todo = input.Todo
	todo.DateTime = input.DateTime
	todo.UserID = input.User.ID

	// TODO 2
	var newTodo, err = s.repository.Save(todo)
	if err != nil {
		return newTodo, err
	}
	
	// TODO 3
	return newTodo, nil

}

func (s *service) UpdateTodo(inputID GetTodoDetailInput, inputData CreateTodoInput) (Todo, error) {
	// TODO 
	// 1. filter Id User, dengan select ke tabel user
	// 2. membandingkan antara id di inputData & todo hasil dari select database, kalau beda return error
	// 3. set data ke struct untuk persiapan update data 
	// 4. panggil Update di package repository dan oper data struct
	// 5. return data struct yg baru di update
	
	// TODO 1
	var todo, err = s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}
	
	// TODO 2
	if todo.UserID != inputData.User.ID {
		return todo, errors.New("todo nya bukan milik si user")
	}
	
	// TODO 3
	todo.UserID = inputData.User.ID
	todo.Todo = inputData.Todo
	todo.DateTime = inputData.DateTime
	
	// TODO 4
	updateTodo, err := s.repository.Update(todo)
	if err != nil {
		return updateTodo, err
	}
	
	// TODO 5
	return updateTodo, nil

}

func (s *service) DeleteTodo(inputID GetTodoDetailInput, inputData CreateTodoInput) (Todo, error) {
	// TODO
	// 1. filter Id User, dengan select ke tabel user
	// 2. membandingkan antara id di inputData & todo hasil dari select database, kalau beda return error
	// 3. panggil method Delete di package repository
	// 4. return data struct todo
	
	// TODO 1 
	var todo, err = s.repository.FindByID(inputID.ID)
	if err != nil {
		return todo, err
	}
	
	// TODO 2
	if todo.UserID != inputData.User.ID {
		return todo, errors.New("todo nya bukan milik si user")
	}
	
	// TODO 3
	var deleteTodo, err2 = s.repository.Delete(todo)
	if err2 != nil {
		return deleteTodo, err
	}
	
	// TODO 4
	return deleteTodo, nil

}
