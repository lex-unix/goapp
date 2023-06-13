package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	Get(id int64) (*User, error)
	Create(u *User) error
	Delete(id int64) error
	User(username string) (*User, error)
}
