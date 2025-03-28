package users

import (
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/internalerrors"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string       `gorm:"primaryKey"`
	Username     string       `gorm:"unique" validate:"required,min=2,max=32"`
	Email        string       `gorm:"unique" validate:"email"`
	Password     string       `validate:"min=8,max=64" gorm:"-" json:"-" graphql:"-"`
	HashPassword string       `gorm:"not null" json:"-" graphql:"-"`
	Todos        []todos.Todo `gorm:"foreignKey:UserID"`
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashPassword), []byte(password))
}

func NewUser(username string, email string, password string) (*User, error) {
	id := xid.New().String()

	user := &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
		Todos:    []todos.Todo{},
	}

	err := internalerrors.ValidateStruct(user)
	if err != nil {
		return nil, err
	}

	err = addHashPassword(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func addHashPassword(user *User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}

	user.HashPassword = string(bytes)
	user.Password = ""

	return nil
}
