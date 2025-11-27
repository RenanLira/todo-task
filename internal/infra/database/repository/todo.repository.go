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
	GetPageInfo(userId string, perPage int, page int) (types.Page, error)
	GetAllByUser(userId string, limit int32, offset int32) ([]*todos.Todo, error)
}

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func (r *TodoRepositoryImpl) Create(todo *todos.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepositoryImpl) Update(todo *todos.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepositoryImpl) Delete(todo *todos.Todo) error {
	return r.db.Delete(todo).Error
}

func (r *TodoRepositoryImpl) Find(id string) (*todos.Todo, error) {
	var todo *todos.Todo
	err := r.db.First(&todo, "id", id).Error

	return todo, err
}

func (r *TodoRepositoryImpl) GetAll(limit int32, offset int32) ([]*todos.Todo, error) {

	var t []*todos.Todo
	err := r.db.Find(&t).Limit(int(limit)).Offset(int(offset)).Error

	return t, err
}

func (r *TodoRepositoryImpl) GetPageInfo(userId string, perPage int, page int) (types.Page, error) {

	var count int64
	r.db.Model(&todos.Todo{}).Where(&todos.Todo{UserID: userId}).Count(&count)

	return types.Page{
		HasNextPage:     count/int64(perPage) > int64(perPage*page),
		HasPreviousPage: page > 1,
		Quantity:        int(count),
	}, nil
}

func (r *TodoRepositoryImpl) GetAllByUser(userId string, limit int32, offset int32) ([]*todos.Todo, error) {
	var todosList []*todos.Todo

	err := r.db.Find(&todosList, &todos.Todo{UserID: userId}).Limit(int(limit)).Offset(int(offset)).Error
	if err != nil {
		return nil, err
	}

	return todosList, nil
}

func NewTodoRepository() *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: database.NewDB()}
}
