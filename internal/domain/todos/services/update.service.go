package services

import (
	"errors"
	"todo-tasks/internal/domain/auth/authorization"
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/utils"
)

func (t *TodoService) UpdateTodo(todoDTO todos.ReqUpdateTodoDTO) (*todos.Todo, error) {
	todo, err := t.TodoRepository.Find(todoDTO.ID)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	err = authorization.UserTodoAllow(todoDTO.UserID, "update", todo.UserID)
	if err != nil {
		return nil, err
	}

	utils.CopyStruct(todo, todoDTO.Fields)

	err = t.TodoRepository.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
