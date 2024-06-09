package auth

import (
	"github.com/dgrijalva/jwt-go"

	"errors"
)

type Service interface {
	GenerateToken(userID int, userName string, userPassword string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("Tfn#@j0lMq2vHte%*&")

func (s *jwtService) GenerateToken(userID int, userName string, userPassword string) (string, error) {
	var claim = jwt.MapClaims{}
	claim["user_id"] = userID
	claim["user_name"] = userName
	claim["user_password"] = userPassword
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	var signedToken, err = token.SignedString(SECRET_KEY)
	if err != nil { 
		return signedToken, err 
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	var token, err = jwt.Parse(encodedToken, func(structToken *jwt.Token) (interface{}, error) {
		_, ok := structToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token tidak lazim")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	
	var token, err = jwt.Parse(encodedToken, func(structToken *jwt.Token) (interface{}, error) {
		_, ok := structToken.Method.(*jwt.SigningMethodHMAC)
		if !ok { 
			return nil, errors.New("token tidak lazim")  
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil { 
		return token, err 
	}
	return token, nil
}
