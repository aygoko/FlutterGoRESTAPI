package repository

import (
	"errors"
)

type User struct {
	data map[string]string
}

func NewUser(repo repository.User) *User {
	return &User{
		repo: repo,
	}
}

func (rs *User) Get(key string) (*string, error) {
	value, exists := rs.data[key]
	if !exists {
		return nil, errors.New("key not found")
	}
	return &value, nil
}

func (rs *User) Put(key string, value string) error {
	rs.data[key] = value
	return nil
}

func (rs *User) Post(key string, value string) error {
	if _, exists := rs.data[key]; exists {
		return errors.New("key already exists")
	}
	rs.data[key] = value
	return nil
}

func (rs *User) Delete(key string) error {
	if _, exists := rs.data[key]; !exists {
		return errors.New("key not found")
	}
	delete(rs.data, key)
	return nil
}
