package services

import "todo-tasks/internal/domain/users"

func (u *UserService) CreateUser(dto users.CreateUserDTO) (*users.User, error) {
	user, err := users.NewUser(dto.Username, dto.Email, dto.Password)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
