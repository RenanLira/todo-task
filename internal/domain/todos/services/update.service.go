package services

import (
	"errors"
	"todo-tasks/internal/domain/auth/authorization"
	"todo-tasks/internal/domain/todos"
)

func (t *TodoService) UpdateTodo(todoDTO todos.Todo) (*todos.Todo, error) {
	todo, err := t.TodoRepository.Find(todoDTO.ID)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	err = authorization.UserTodoAllow(todoDTO.UserID, "update", todo.UserID)
	if err != nil {
		return nil, err
	}

	err = t.TodoRepository.Update(&todos.Todo{
		ID:     todo.ID,
		Done:   todoDTO.Done,
		Text:   todoDTO.Text,
		UserID: todo.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &todoDTO, nil
}
