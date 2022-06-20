package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUserInput(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

//s berisi repository, dan Interface bisa panggil service
func (s *service) RegisterUserInput(input RegisterUserInput) (User, error) {
	user := User{}
	user.Username = input.Username
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
