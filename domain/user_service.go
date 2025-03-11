package repository

type UserService interface {
	Save(user *User) (*User, error)
	Get(login string) (*User, error)
	GetByEmail(email string) (*User, error)
	List() ([]*User, error)
	Delete(login string) error
}
