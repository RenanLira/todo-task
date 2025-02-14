package users

import "github.com/rs/xid"

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
}

func NewUser(username string, email string) *User {
	id := xid.New().String()

	return &User{
		ID:       id,
		Username: username,
		Email:    email,
	}
}
