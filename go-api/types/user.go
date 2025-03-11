package types

import (
	"errors"

	"github.com/aygoko/FlutterGoRESTAPI/repository"
	"github.com/aygoko/FlutterGoRESTAPI/usecases/service"
	"golang.org/x/crypto/bcrypt"
)

func (s *repository.UserService) CreateUser(login, email, password string) (*repository.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if login == "" || email == "" {
		return nil, errors.New("login and email are required")
	}

	if _, exists := s.Users[login]; exists {
		return nil, errors.New("user with this login already exists")
	}

	if existingLogin, exists := s.Emails[email]; exists {
		return nil, errors.New("email already in use by " + existingLogin)
	}

	user := &repository.User{
		ID:       service.GenerateUserID(),
		Login:    login,
		Email:    email,
		Password: string(hashedPassword),
	}

	s.Users[login] = user
	s.Emails[email] = login

	return user, nil
}

func (s *repository.UserService) GetUserByLogin(login string) (*repository.User, error) {
	user, exists := s.Users[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *repository.UserService) GetUserByEmail(email string) (*repository.User, error) {
	login, exists := s.Emails[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	user, exists := s.Users[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *repository.UserService) GetAllUsers() []*repository.User {
	users := make([]*repository.User, 0, len(s.Users))
	for _, user := range s.Users {
		users = append(users, user)
	}
	return users
}

func (s *repository.UserService) DeleteUser(login string) error {
	user, exists := s.Users[login]
	if !exists {
		return errors.New("user not found")
	}

	delete(s.Users, login)
	delete(s.Emails, user.Email)

	return nil
}
