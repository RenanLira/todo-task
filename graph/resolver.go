package graph

import (
	todo_service "todo-tasks/internal/domain/todos/services"
	user_service "todo-tasks/internal/domain/users/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoService  *todo_service.TodoService
	UserResolver *user_service.UserService
}
