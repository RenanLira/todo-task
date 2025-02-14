package repository

import (
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/infra/database"

	"gorm.io/gorm"
)

type TodoRepository interface {
	database.Repository[todos.Todo]
	GetAllTodos() ([]*todos.Todo, error)
}

type TodoRepositoryImp struct {
	db *gorm.DB
}

func (r *TodoRepositoryImp) Create(todo *todos.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepositoryImp) Update(todo *todos.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepositoryImp) Delete(todo *todos.Todo) error {
	return r.db.Delete(todo).Error
}

func (r *TodoRepositoryImp) Find(id uint) (*todos.Todo, error) {
	var todo todos.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepositoryImp) GetAllTodos() ([]*todos.Todo, error) {
	var todos []*todos.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func NewTodoRepository() *TodoRepositoryImp {
	return &TodoRepositoryImp{db: database.NewDB()}
}
