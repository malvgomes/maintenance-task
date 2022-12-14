package service

import (
	"context"
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/repository"
)

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{repository: ctx.Value("userRepository").(repository.UserRepository)}
}

type GetUserService struct {
	repository repository.UserRepository
}

func (u *GetUserService) GetUser(username, password string) (*model.User, error) {
	return u.repository.GetUser(username, password)
}

func (u *GetUserService) GetUserByID(ID int) (*model.User, error) {
	return u.repository.GetUserByID(ID)
}

func (u *GetUserService) GetUsersByRole(role model.Role) ([]*model.User, error) {
	return u.repository.GetUsersByRole(role)
}
