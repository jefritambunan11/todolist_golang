package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUserByID(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	var user User
	user.Name = input.Name
	user.Email = input.Email

	var passwordHash, err = bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil { return user, err }

	user.Password = string(passwordHash)

	var newUser, err2 = s.repository.Save(user)
	if err2 != nil { return newUser, err2 }

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	var email = input.Email
	var password = input.Password

	var user, err = s.repository.FindByEmail(email)
	if err != nil { return user, err }

	if user.ID == 0 {
		return user, errors.New("User Tidak Ditemukan")
	}

	var _err_ = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if _err_ != nil { return user, errors.New("password salah") }

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	var email = input.Email

	var user, err = s.repository.FindByEmail(email)
	if err != nil { return false, err }

	if user.ID > 0 { return false, nil}

	return true, nil
}

func (s *service) GetUserByID(id int) (User, error) {
	var user, err = s.repository.FindByID(id)
	
	if err != nil { return user, err }
	
	if user.ID == 0 {
		return user, errors.New("Tidak Ada User Menggunakan ID Tersebut")
	}

	return user, nil
}
