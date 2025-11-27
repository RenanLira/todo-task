package database

import (
	"fmt"
	"os"
	"sync"
	"time"
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/domain/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	Find(id string) (*T, error)
}

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func NewDB() *gorm.DB {

	dbOnce.Do(func() {
		var info = map[string]string{
			"host":     os.Getenv("DB_HOST"),
			"port":     os.Getenv("DB_PORT"),
			"user":     os.Getenv("DB_USERNAME"),
			"password": os.Getenv("DB_PASSWORD"),
			"dbname":   os.Getenv("DB_DATABASE"),
			"sslmode":  "disable",
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			info["host"], info["user"], info["password"], info["dbname"], info["port"],
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&users.User{}, &todos.Todo{})

		dbDB, _ := db.DB()
		dbDB.SetMaxIdleConns(10)
		dbDB.SetMaxOpenConns(100)
		dbDB.SetConnMaxLifetime(time.Hour)

		dbInstance = db
	})

	return dbInstance
}
