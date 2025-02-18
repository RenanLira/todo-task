package services

import "errors"

func (t *TodoService) DeleteTodo(id string) error {
	todo, _ := t.TodoRepository.Find(id)
	if todo == nil {
		return errors.New("todo not found")
	}

	err := t.TodoRepository.Delete(todo)
	if err != nil {
		return err
	}
	
	return nil
}
