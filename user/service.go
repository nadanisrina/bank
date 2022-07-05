package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUserInput(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	CheckEmail(input CheckEmailInput) (bool, error)
	UploadAvatar(id int, fileLocation string) (User, error)
	GetUserById(id int) (User, error)
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

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckEmail(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	fmt.Println("user", user)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil

}

func (s *service) UploadAvatar(id int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation
	updateUser, err := s.repository.Update(user)
	if err != nil {
		return user, nil
	}
	return updateUser, err
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found with id given")
	}
	return user, nil
}
