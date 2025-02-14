package services

import (
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/infra/database/repository"
)

type TodoService struct {
	TodoRepository repository.TodoRepository
}

func (s *TodoService) CreateTodo(dto todos.CreateTodoDTO) (*todos.Todo, error) {
	todo := todos.NewTodo(dto.Text)

	err := s.TodoRepository.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) GetAllTodos(dto todos.ReqGetAllTodosDTO) ([]*todos.Todo, error) {
	
	todos, err := s.TodoRepository.GetAllTodos()
	
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func NewTodoService() *TodoService {
	return &TodoService{
		TodoRepository: repository.NewTodoRepository(),
	}
}
