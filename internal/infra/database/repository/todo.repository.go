package repository

import (
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/infra/database"
	"todo-tasks/internal/utils/types"

	"gorm.io/gorm"
)

type TodoRepository interface {
	database.Repository[todos.Todo]
	GetAll(limit int32, offset int32) ([]*todos.Todo, error)
	GetPageInfo(perPage int, page int) (types.Page, error)
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

func (r *TodoRepositoryImp) Find(id string) (*todos.Todo, error) {
	var todo *todos.Todo
	err := r.db.First(&todo, "id", id).Error

	return todo, err
}

func (r *TodoRepositoryImp) GetAll(limit int32, offset int32) ([]*todos.Todo, error) {

	var todos []*todos.Todo
	err := r.db.Find(&todos).Limit(int(limit)).Offset(int(offset)).Error

	return todos, err
}

func (r *TodoRepositoryImp) GetPageInfo(perPage int, page int) (types.Page, error) {

	var count int64
	r.db.Model(&todos.Todo{}).Count(&count)

	return types.Page{
		HasNextPage:     count/int64(perPage) > int64(perPage*page),
		HasPreviousPage: page > 1,
		Quantity:        int(count),
	}, nil
}

func NewTodoRepository() *TodoRepositoryImp {
	return &TodoRepositoryImp{db: database.NewDB()}
}
