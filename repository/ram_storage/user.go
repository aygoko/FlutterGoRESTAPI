package repository

import (
	"errors"

	"github.com/aygoko/FlutterGoRESTAPI/repository"
)

type User struct {
	data map[string]*repository.UserService
}

func NewUser() *repository.UserService {
	return &repository.UserService{
		data: make(map[string]*repository.UserService),
	}
}

func (rs *User) Save(user *User) error {
	if _, exists := rs.data[user.Login]; exists {
		return errors.New("user already exists")
	}
	rs.data[user.Login] = user
	return nil
}

func (rs *User) Get(login string) (*User, error) {
	user, exists := rs.data[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (rs *User) GetByEmail(email string) (*User, error) {
	for _, user := range rs.data {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// List returns all users
func (rs *User) List() ([]*User, error) {
	users := make([]*User, 0, len(rs.data))
	for _, user := range rs.data {
		users = append(users, user)
	}
	return users, nil
}

// Delete removes a user by login
func (rs *User) Delete(login string) error {
	if _, exists := rs.data[login]; !exists {
		return errors.New("user not found")
	}
	delete(rs.data, login)
	return nil
}
