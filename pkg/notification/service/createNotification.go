package service

import (
	"context"
	"maintenance-task/pkg/notification/model"
	"maintenance-task/pkg/notification/repository"
)

func NewCreateNotificationService(ctx context.Context) *CreateNotificationService {
	return &CreateNotificationService{
		repository: ctx.Value("notificationRepository").(repository.NotificationRepository)}
}

type CreateNotificationService struct {
	repository repository.NotificationRepository
}

func (u *CreateNotificationService) CreateNotification(input model.CreateNotification) error {
	return u.repository.CreateNotification(input)
}
