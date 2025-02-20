package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"
	"todo-tasks/internal/domain/auth/types"
	"todo-tasks/internal/domain/users"
	"todo-tasks/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

func validateToken(tokenString string) (*users.User, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT_TOKEN")), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims, _ := token.Claims.(types.JwtClaims)

	return &claims.User, nil
}

func AuthMiddleware() func(next http.Handler) http.Handler {
	// repo := repository.NewUserRepository()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				utils.SendJSON(w, map[string]string{"error": "Invalid Authorization header format"}, http.StatusUnauthorized)
				return
			}

			token := tokenParts[1]

			user, err := validateToken(token)
			if err != nil {
				utils.SendJSON(w, map[string]string{"error": "Invalid token", "message": err.Error()}, http.StatusUnauthorized)
				return
			}

			// Do something
			ctx := context.WithValue(r.Context(), types.UserCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
