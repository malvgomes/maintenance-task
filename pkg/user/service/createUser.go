package service

import (
	"context"
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/repository"
)

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{repository: ctx.Value("userRepository").(repository.UserRepository)}
}

type CreateUserService struct {
	repository repository.UserRepository
}

func (u *CreateUserService) CreateUser(user model.CreateUser) (int, error) {
	return u.repository.CreateUser(user)
}
