package todos

import "github.com/rs/xid"

type Todo struct {
	ID   string `gorm:"primaryKey"`
	Text string `gorm:"not null"`
	Done bool `gorm:"default:false"`
}

func NewTodo(text string) *Todo {
	id := xid.New()

	return &Todo{
		ID:   id.String(),
		Text: text,
		Done: false,
	}
}
