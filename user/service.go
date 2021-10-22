package user 

import(
	"golang.org/x/crypto/bcrypt"

	"errors"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
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
	user.Occupation = input.Occupation 
	
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost) 
	if err != nil {
		return user, err
	}
	
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	
	return newUser, nil	
}

func (s *service) Login(input LoginInput)(User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	_err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) 
	if _err != nil {
	//	return user, _err
		return user, errors.New("Wrong Password")
	}

	return user, nil
}