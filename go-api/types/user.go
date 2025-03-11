package types

import (
	"errors"

	"github.com/aygoko/FlutterGoRESTAPI/domain"
	"github.com/aygoko/FlutterGoRESTAPI/usecases/service"
	"golang.org/x/crypto/bcrypt"
)

// UserService should be in the service package, so adjust receiver types
// @Summary Create a new user
// @Description Creates a user with login, email, and hashed password
// @Tags Users
// @Param login formData string true "User login"
// @Param email formData string true "User email"
// @Param password formData string true "User password (plaintext)"
// @Success 201 {object} types.User
// @Failure 400 {string} error "Invalid request or duplicate user"
// @Router /api/users [post]
func (s *service.UserService) CreateUser(login, email, password string) (*domain.User, error) { // Use types.User here
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if login == "" || email == "" {
		return nil, errors.New("login and email are required")
	}

	// Use exported fields from service.UserService (uppercase)
	if _, exists := s.Users[login]; exists {
		return nil, errors.New("user with this login already exists")
	}

	if existingLogin, exists := s.Emails[email]; exists {
		return nil, errors.New("email already in use by " + existingLogin)
	}

	user := &User{ // Create instance of types.User
		ID:       service.GenerateUserID(),
		Login:    login,
		Email:    email,
		Password: string(hashedPassword),
	}

	s.Users[login] = user // Ensure service.UserService has Users map
	s.Emails[email] = login

	return user, nil
}

// @Summary Get a user by login
// @Description Retrieves a user by their login
// @Tags Users
// @Param login path string true "User login"
// @Success 200 {object} types.User
// @Failure 404 {string} error "User not found"
// @Router /api/users/{login} [get]
func (s *service.UserService) GetUserByLogin(login string) (*domain.User, error) { // Return types.User
	user, exists := s.Users[login] // Access exported Users map
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil // Directly return the pointer
}

// @Summary Get a user by email
// @Description Retrieves a user by their email
// @Tags Users
// @Param email path string true "User email" // Changed to path
// @Success 200 {object} types.User
// @Failure 404 {string} error "User not found"
// @Router /api/users/email/{email} [get]
func (s *service.UserService) GetUserByEmail(email string) (*User, error) {
	login, exists := s.Emails[email] // Use exported Emails map
	if !exists {
		return nil, errors.New("user not found")
	}
	user, exists := s.Users[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// @Summary List all users
// @Description Returns a list of all registered users
// @Tags Users
// @Success 200 {array} types.User
// @Router /api/users [get]
func (s *service.UserService) GetAllUsers() []*User {
	users := make([]*User, 0, len(s.Users))
	for _, user := range s.Users {
		users = append(users, user)
	}
	return users
}

// @Summary Delete a user by login
// @Description Removes a user by their login
// @Tags Users
// @Param login path string true "User login to delete"
// @Success 204 "User deleted successfully"
// @Failure 404 {string} error "User not found"
// @Router /api/users/{login} [delete]
func (s *service.UserService) DeleteUser(login string) error {
	user, exists := s.Users[login]
	if !exists {
		return errors.New("user not found")
	}

	delete(s.Users, login)
	delete(s.Emails, user.Email) // Ensure user.Email is accessible

	return nil
}
