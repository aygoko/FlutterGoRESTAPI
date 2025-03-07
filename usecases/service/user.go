package service

import (
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func generateUserID() string {
	return uuid.New().String()
}
