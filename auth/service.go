package auth

import (
	"github.com/dgrijalva/jwt-go"

	"errors"
)

type Service interface {
	GenerateToken(userID int, userName string, userPassword string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("Tfn#@j0lMq2vHte%*&")

func (s *jwtService) GenerateToken(userID int, userName string, userPassword string) (string, error) {
	// TODO
	// 1. pake struct jwt dan dimana field struct nya di set user_id, user_name, user_password dan  datanya berasal dari params
	// 2. bagian ini enkrip struct claim dengan pola JWT dan return nya struct Token
	// 3. bagian ini gabungkan antar hasil struct claim dan string SECRET_KEY
	// 4. bagian ini return hasil token akhir

	// TODO 1
	var claim = jwt.MapClaims{}
	claim["user_id"] = userID
	claim["user_name"] = userName
	claim["user_password"] = userPassword

	// TODO 2
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// TODO 3
	var signedToken, err = token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	// TODO 4
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	// TODO
	// 1. bagian ini memeriksa kode token, params 1 -> kode token, params 2 -> satu fungsi anonymous dan menggunakan 1 params struct Token dari package jwt
	// 2. cek apakah saat proses terjemahkan ada error
	// 3. return kan nilai token yg artinya token valid

	// TODO 1
	var token, err = jwt.Parse(encodedToken, func(structToken *jwt.Token) (interface{}, error) {
		_, ok := structToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("token tidak lazim")
		}

		return []byte(SECRET_KEY), nil
	})

	// TODO 2
	if err != nil {
		return token, err
	}

	// TODO 3
	return token, nil
}
