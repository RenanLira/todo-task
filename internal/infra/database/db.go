package database

import (
	"fmt"
	"os"
	"todo-tasks/internal/domain/todos"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	Find(id string) (*T, error)
}

func NewDB() *gorm.DB {
	var info = map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"user":     os.Getenv("DB_USERNAME"),
		"password": os.Getenv("DB_PASSWORD"),
		"dbname":   os.Getenv("DB_DATABASE"),
		"sslmode":  "disable",
	}

	var dsn string
	for key, value := range info {
		dsn += fmt.Sprintf("%s=%s ", key, value)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&todos.Todo{})

	return db
}
