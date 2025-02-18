package directives

import (
	"context"
	"errors"
	"todo-tasks/internal/domain/auth/decorators"
	"todo-tasks/internal/domain/users"

	"github.com/99designs/gqlgen/graphql"
)

func AuthDirective(ctx context.Context, obj any, next graphql.Resolver) (any, error) {

	user := ctx.Value(decorators.UserCtxKey)

	switch user.(type) {
	case *users.User:
		return next(ctx)
	default:
		return nil, errors.New("Unauthorized")
	}

}
