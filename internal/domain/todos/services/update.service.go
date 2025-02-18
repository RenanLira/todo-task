package services

import (
	"errors"
	"todo-tasks/internal/domain/todos"
)

func (t *TodoService) UpdateTodo(todo todos.Todo) (*todos.Todo, error) {

	if t, _ := t.TodoRepository.Find(todo.ID); t == nil {
		return nil, errors.New("todo not found")
	}

	err := t.TodoRepository.Update(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
