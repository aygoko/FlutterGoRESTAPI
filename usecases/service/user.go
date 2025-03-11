package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	repository "github.com/aygoko/FlutterGoRESTAPI/repository/ram_storage"

	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserService struct {
	repo   repository.User
	users  map[string]*User
	emails map[string]string
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo:   repo,
		users:  make(map[string]*User),
		emails: make(map[string]string),
	}
}

func (s *UserService) CreateUser(login, email, password string) (*User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

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
		ID:       generateUserID(),
		Login:    login,
		Email:    email,
		Password: string(hashedPassword),
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

func (s *UserService) GetAllUsers() []*User {
	var users []*User
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *UserService) DeleteUser(login string) error {
	user, exists := s.users[login]
	if !exists {
		return errors.New("user not found")
	}

	delete(s.users, login)
	delete(s.emails, user.Email)

	return nil
}

func generateUserID() string {
	return uuid.New().String()
}
