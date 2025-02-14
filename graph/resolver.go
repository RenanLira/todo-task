package graph

import (
	"todo-tasks/internal/domain/todos/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoService *services.TodoService
}
