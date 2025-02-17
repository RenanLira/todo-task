package todos

import (
	"todo-tasks/internal/internalerrors"

	"github.com/rs/xid"
)

type Todo struct {
	ID   string `gorm:"primaryKey"`
	Text string `gorm:"not null" validate:"required,min=4,max=255"`
	Done bool   `gorm:"default:false"`
	UserID string `gorm:"not null"`
}

func NewTodo(text string, userID string) (*Todo, error) {
	id := xid.New()

	todo := &Todo{
		ID:   id.String(),
		Text: text,
		Done: false,
		UserID: userID,
	}

	err := internalerrors.ValidateStruct(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
