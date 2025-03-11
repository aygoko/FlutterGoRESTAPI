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

// @Summary Create a new user
// @Description Creates a user with login, email, and hashed password
// @Tags Users
// @Param login query string true "User login"
// @Param email query string true "User email"
// @Param password query string true "User password (plaintext)"
// @Success 201 {object} User
// @Failure 400 {string} error "Invalid request or duplicate user"
// @Router /api/users [post]
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

// @Summary Get a user by login
// @Description Retrieves a user by their login
// @Tags Users
// @Param login path string true "User login"
// @Success 200 {object} User
// @Failure 404 {string} error "User not found"
// @Router /api/users/{login} [get]
func (s *UserService) GetUserByLogin(login string) (*User, error) {
	user, exists := s.users[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// @Summary Get a user by email
// @Description Retrieves a user by their email
// @Tags Users
// @Param email query string true "User email"
// @Success 200 {object} User
// @Failure 404 {string} error "User not found"
// @Router /api/users/email/{email} [get]
func (s *UserService) GetUserByEmail(email string) (*User, error) {
	login, exists := s.emails[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return s.users[login], nil
}

// @Summary List all users
// @Description Returns a list of all registered users
// @Tags Users
// @Success 200 {array} User
// @Router /api/users [get]
func (s *UserService) GetAllUsers() []*User {
	var users []*User
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// @Summary Delete a user by login
// @Description Removes a user by their login
// @Tags Users
// @Param login patstring true "User login to delete"
// @Success 204 "User deleted successfully"
// @Failure 404 {string} error "User not found"
// @Router /api/users/{login} [delete]
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
