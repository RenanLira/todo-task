package services

import (
	"errors"
	"todo-tasks/internal/domain/auth/authorization"
	"todo-tasks/internal/domain/todos"
)

func (t *TodoService) DeleteTodo(id string, userId string) (*todos.Todo, error) {
	todo, _ := t.TodoRepository.Find(id)
	if todo == nil {
		return nil, errors.New("todo not found")
	}

	err := authorization.UserTodoAllow(userId, "delete", todo.UserID)
	if err != nil {
		return nil, err
	}

	err = t.TodoRepository.Delete(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
