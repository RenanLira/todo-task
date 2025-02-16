package services

import "todo-tasks/internal/domain/todos"

func (s *TodoService) CreateTodo(dto todos.CreateTodoDTO) (*todos.Todo, error) {
	todo := todos.NewTodo(dto.Text)

	err := s.TodoRepository.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
