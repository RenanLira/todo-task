package services

import "todo-tasks/internal/domain/todos"

func (t *TodoService) CreateTodo(dto todos.CreateTodoDTO) (*todos.Todo, error) {
	todo, err := todos.NewTodo(dto.Text, dto.UserID)
	if err != nil {
		return nil, err
	}

	err = t.TodoRepository.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
