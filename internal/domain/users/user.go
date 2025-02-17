package users

import (
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/internalerrors"

	"github.com/rs/xid"
)

type User struct {
	ID       string       `gorm:"primaryKey"`
	Username string       `gorm:"unique"`
	Email    string       `gorm:"unique" validate:"email"`
	Todos    []todos.Todo `gorm:"foreignKey:UserID"`
}

func NewUser(username string, email string) (*User, error) {
	id := xid.New().String()

	user := &User{
		ID:       id,
		Username: username,
		Email:    email,
		Todos:    []todos.Todo{},
	}

	err := internalerrors.ValidateStruct(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
