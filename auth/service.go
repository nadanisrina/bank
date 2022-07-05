package auth

import "github.com/golang-jwt/jwt/v4"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("s3cr3t_k3y")

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
