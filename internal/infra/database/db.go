package database

import (
	"todo-tasks/internal/domain/todos"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	Find(id uint) (*T, error)
}

const dsn = "host=localhost user=postgres password=postgres dbname=tododb port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

func NewDB() (*gorm.DB, error) {
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	
	if err != nil {
		return nil, err
	}
	
	db.AutoMigrate(&todos.Todo{})
	
	return db, nil
}
