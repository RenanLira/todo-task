package services

import "todo-tasks/internal/domain/todos"

func (t *TodoService) Find(id string) *todos.Todo {

	todo, err := t.TodoRepository.Find(id)
	if err != nil || todo == nil {
		return nil
	}

	return todo

}
