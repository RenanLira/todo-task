package types

import (
	"github.com/golang-jwt/jwt/v5"
	"todo-tasks/internal/domain/users"
)

type JwtClaims struct {
	users.User
	jwt.RegisteredClaims
}
