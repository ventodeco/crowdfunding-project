package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(request RegisterUserRequest) (User, error)
	LoginUser(request LoginRequest) (User, error)
	IsEmailAvailable(request CheckEmailRequest) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(request RegisterUserRequest) (User, error) {
	user := User{}
	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

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

func (s *service) LoginUser(request LoginRequest) (User, error) {
	email := request.Email
	password := request.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(request CheckEmailRequest) (bool, error) {
	email := request.Email

	userByEmail, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if userByEmail.ID == 0 {
		return true, nil
	}

	return false, nil
}
