package repository

import (
	"github.com/aygoko/FlutterGoRESTAPI/usecases/service"
)

type UserService interface {
	Save(user *service.User) error
	Get(login string) (*service.User, error)
	GetByEmail(email string) (*service.User, error)
	List() ([]*service.User, error)
	Delete(login string) error
}
