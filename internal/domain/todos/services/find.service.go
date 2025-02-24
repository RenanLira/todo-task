package services

import (
	"fmt"
	"todo-tasks/internal/domain/auth/authorization"
	"todo-tasks/internal/domain/todos"
)

func (t *TodoService) Find(todoId, userId string) (*todos.Todo, error) {

	todo, err := t.TodoRepository.Find(todoId)
	if err != nil || todo == nil {
		return nil, fmt.Errorf("not found")
	}

	err = authorization.UserTodoAllow(userId, "get", todo.UserID)
	if err != nil {
		return nil, err
	}

	return todo, nil

}
