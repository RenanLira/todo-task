package types

import (
	"github.com/golang-jwt/jwt/v5"
	"todo-tasks/internal/domain/users"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	User users.User `json:"user"`
}
