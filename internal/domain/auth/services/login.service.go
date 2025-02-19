package services

import (
	"os"
	"time"
	"todo-tasks/internal/domain/users"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	users.User `json:"user"`
	jwt.RegisteredClaims
}

func generateToken(user *users.User) (string, error) {
	claims := &CustomClaims{
		*user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			ID: user.ID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT_TOKEN")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthService) Login(email string, password string) (*users.User, string, error) {
	user, err := a.repo.FindByEmail(email)
	if err != nil {
		return nil, "", err
	}

	err = user.ComparePassword(password)
	if err != nil {
		return nil, "", err
	}

	tokenString, err := generateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}
