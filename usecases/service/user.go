package service

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

type UserService struct {
	users  map[string]*User
	emails map[string]string
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[string]*User),
		emails: make(map[string]string),
	}
}

func (s *UserService) CreateUser(login, email, password string) (*User, error) {

	if login == "" || email == "" {
		return nil, errors.New("login and email are required")
	}

	if _, exists := s.users[login]; exists {
		return nil, errors.New("user with this login already exists")
	}

	if existingLogin, exists := s.emails[email]; exists {
		return nil, errors.New("email already in use by " + existingLogin)
	}

	user := &User{
		ID:    generateUserID(),
		Login: login,
		Email: email,
	}

	s.users[login] = user
	s.emails[email] = login

	return user, nil
}

func (s *UserService) GetUserByLogin(login string) (*User, error) {
	user, exists := s.users[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	login, exists := s.emails[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return s.users[login], nil
}

func generateUserID() string {
	return uuid.New().String()
}
