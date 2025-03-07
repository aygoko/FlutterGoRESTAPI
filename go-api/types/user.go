package types

import (
	"errors"
	"service"
)

func (u *service.User) CreateUser(login, password string) (*service.User, error) {
	if _, exists := u.users[login]; exists {
		return nil, errors.New("user already exists")
	}
	user := &service.User{
		ID:       service.generateUserID(),
		Login:    login,
		Password: password,
	}
	u.users[login] = user
	return user, nil
}

func (u *service.User) GetUserByLogin(login string) (*service.User, error) {
	user, exists := u.users[login]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
