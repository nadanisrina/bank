package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var envSecretKey = os.Getenv("SECRET_KEY")
var SECRET_KEY = []byte(envSecretKey)

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedKey, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {

		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
