package service

import (
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
	Repo repository.UserService
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func GenerateUserID() string {
	return uuid.New().String()
}
