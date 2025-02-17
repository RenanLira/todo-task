package middlewares

import (
	"net/http"
)

func AuthMiddleware() func(next http.Handler) http.Handler {
	// repo := repository.NewUserRepository()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			id, err := r.Cookie("id")
			if err != nil || id == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return

			}

			// Do something
			next.ServeHTTP(w, r)
		})
	}

}
