package service

import (
	"context"
	"maintenance-task/pkg/user/repository"
)

func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{repository: ctx.Value("userRepository").(repository.UserRepository)}
}

type DeleteUserService struct {
	repository repository.UserRepository
}

func (u *DeleteUserService) DeleteUser(username string) error {
	return u.repository.DeleteUser(username)
}
