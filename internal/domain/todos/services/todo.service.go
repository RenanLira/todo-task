package services

import (
	"todo-tasks/internal/infra/database/repository"
)

type TodoService struct {
	TodoRepository repository.TodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{
		TodoRepository: repository.NewTodoRepository(),
	}
}
