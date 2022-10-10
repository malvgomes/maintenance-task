package repository

import (
	"context"
	"maintenance-task/pkg/user/model"
)

type UserRepository interface {
	GetUser(username, password string) (*model.User, error)
	GetUserByID(ID int) (*model.User, error)
	GetUsersByRole(role model.Role) ([]*model.User, error)
	CreateUser(user model.CreateUser) (int, error)
	DeleteUser(userID int) error
}

func GetUserRepository(ctx context.Context) UserRepository {
	return NewUserRepositoryMariaDB(ctx)
}
