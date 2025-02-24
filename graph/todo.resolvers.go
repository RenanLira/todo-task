package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"todo-tasks/graph/generated"
	"todo-tasks/graph/model"
	"todo-tasks/internal/domain/auth/types"
	"todo-tasks/internal/domain/todos"
	"todo-tasks/internal/domain/users"
	"todo-tasks/internal/utils"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*todos.Todo, error) {
	user := ctx.Value(types.UserCtxKey).(*users.User)

	todo, err := r.TodoService.CreateTodo(todos.CreateTodoDTO{
		Text:   input.Text,
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, update model.UpdateTodo) (*todos.Todo, error) {
	user := ctx.Value(types.UserCtxKey).(*users.User)

	todo, err := r.TodoService.UpdateTodo(todos.ReqUpdateTodoDTO{
		ID:     id,
		UserID: user.ID,
		Fields: todos.UpdateFields{
			Text: update.Text,
			Done: update.Done,
		},
	})

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*todos.Todo, error) {
	user := ctx.Value(types.UserCtxKey).(*users.User)

	todo, err := r.TodoService.DeleteTodo(id, user.ID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// MyTodos is the resolver for the myTodos field.
func (r *queryResolver) MyTodos(ctx context.Context, page *model.PageInput) (*model.TodosResponse, error) {
	user := ctx.Value(types.UserCtxKey).(*users.User)

	var req todos.ReqGetAllTodosDTO
	utils.CopyStruct(&req, *page)

	res, err := r.TodoService.GetAllTodos(req, user.ID)
	if err != nil {
		return nil, err
	}

	return &model.TodosResponse{
		Todos: res.Todos,
		PageInfo: &model.PageInfo{
			HasNextPage:     res.Page.HasNextPage,
			HasPreviousPage: res.Page.HasPreviousPage,
			Quantity:        int32(res.Page.Quantity),
		},
	}, nil
}

// TodoByID is the resolver for the todoById field.
func (r *queryResolver) TodoByID(ctx context.Context, id string) (*todos.Todo, error) {
	user := ctx.Value(types.UserCtxKey).(*users.User)

	todo, err := r.TodoService.Find(id, user.ID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *queryResolver) Todos(_ context.Context, _ *model.PageInput) (*model.TodosResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
*/
