package services

import "todo-tasks/internal/infra/database/repository"

type UserService struct {
	repo repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewUserRepository(),
	}
}
