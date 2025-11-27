package repository

import (
	"todo-tasks/internal/domain/users"
	"todo-tasks/internal/infra/database"

	"gorm.io/gorm"
)

type UserRepository interface {
	database.Repository[users.User]
	FindByEmail(email string) (*users.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) Create(user *users.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepositoryImpl) Update(user *users.User) error {
	return u.db.Save(user).Error
}

func (u *UserRepositoryImpl) Delete(user *users.User) error {
	return u.db.Delete(user).Error
}

func (u *UserRepositoryImpl) Find(id string) (*users.User, error) {
	var user users.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*users.User, error) {
	var user users.User
	err := u.db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{db: database.NewDB()}
}
