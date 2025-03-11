package ram_storage

import (
	"errors"

	repository "github.com/aygoko/FlutterGoRESTAPI/domain" // Import domain package
)

func NewUserRepository() repository.UserService { // Returns the interface
	return &UserRepositoryRAM{
		data: make(map[string]*repository.User),
	}
}

// UserRepositoryRAM implements the UserService interface
type UserRepositoryRAM struct {
	data map[string]*repository.User // Use domain.User
}

// Save stores a user
func (r *UserRepositoryRAM) Save(user *repository.User) (*repository.User, error) {
	if _, exists := r.data[user.Login]; exists {
		return nil, errors.New("user already exists")
	}
	r.data[user.Login] = user
	return user, nil // Return user and error
}

// Get retrieves a user by login
func (r *UserRepositoryRAM) Get(login string) (*repository.User, error) {
	user, exists := r.data[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetByEmail searches for a user by email
func (r *UserRepositoryRAM) GetByEmail(email string) (*repository.User, error) {
	for _, user := range r.data {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// List returns all users
func (r *UserRepositoryRAM) List() ([]*repository.User, error) {
	users := make([]*repository.User, 0, len(r.data))
	for _, user := range r.data {
		users = append(users, user)
	}
	return users, nil
}

// Delete removes a user by login
func (r *UserRepositoryRAM) Delete(login string) error {
	if _, exists := r.data[login]; !exists {
		return errors.New("user not found")
	}
	delete(r.data, login)
	return nil
}
