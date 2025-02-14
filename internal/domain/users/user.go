package users

type User struct {
	ID       string
	Username string
	Email    string
}

func NewUser() *User {
	return &User{}
}
