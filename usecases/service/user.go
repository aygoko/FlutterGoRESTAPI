package service

import (
	repository "github.com/aygoko/FlutterGoRESTAPI/domain"
)

// UserService holds the repository instance
type UserService struct {
	Repo repository.UserService
}

// NewUserService creates a new service instance
func NewUserService(repo repository.UserService) *UserService {
	return &UserService{
		Repo: repo,
	}
}

// Save delegates to repository
// usecases/service/user.go
func (s *UserService) Save(user *repository.User) (*repository.User, error) {
	return s.Repo.Save(user)
}

// Get delegates to repository
func (s *UserService) Get(login string) (*repository.User, error) {
	return s.Repo.Get(login)
}

// GetByEmail delegates to repository
func (s *UserService) GetByEmail(email string) (*repository.User, error) {
	return s.Repo.GetByEmail(email)
}

// List delegates to repository
func (s *UserService) List() ([]*repository.User, error) {
	return s.Repo.List()
}

// Delete delegates to repository
func (s *UserService) Delete(login string) error {
	return s.Repo.Delete(login)
}
