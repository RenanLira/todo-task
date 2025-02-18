package services

import "todo-tasks/internal/infra/database/repository"

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: repository.NewUserRepository(),
	}
}
