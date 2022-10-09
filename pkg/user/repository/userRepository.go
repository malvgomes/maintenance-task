package repository

import (
	"context"
	"maintenance-task/pkg/user/model"
)

type UserRepository interface {
	CreateUser(user model.CreateUser) error
}

func GetUserRepository(ctx context.Context) UserRepository {
	return NewUserRepositoryMariaDB(ctx)
}
