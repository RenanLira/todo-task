package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"
	"todo-tasks/internal/domain/auth/authorization"
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

	claims := token.Claims.(jwt.MapClaims)
	var user users.User

	err = utils.MapToStruct(claims["user"].(map[string]any), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func setEnforcerGroupPolicy(user *users.User) error {
	e := authorization.GetEnforcer()
	_, err := e.AddGroupingPolicy(user.ID, "user")
	return err
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

			err = setEnforcerGroupPolicy(user)
			if err != nil {
				utils.SendJSON(w, map[string]string{"error": "Failed to add grouping policy", "message": err.Error()}, http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), types.UserCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}
