package repository

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
