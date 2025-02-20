package graph

import (
	authservice "todo-tasks/internal/domain/auth/services"
	todoservice "todo-tasks/internal/domain/todos/services"
	userservice "todo-tasks/internal/domain/users/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoService  *todoservice.TodoService
	UserResolver *userservice.UserService
	AuthResolver *authservice.AuthService
}
