package services

import "todo-tasks/internal/domain/users"

func (u *UserService) CreateUser(username string, email string) (*users.User, error) {
	user, err := users.NewUser(username, email)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
