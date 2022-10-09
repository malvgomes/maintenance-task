package service

import (
	"context"
	"maintenance-task/pkg/notification/repository"
)

func NewDeleteNotificationService(ctx context.Context) *DeleteNotificationService {
	return &DeleteNotificationService{
		repository: ctx.Value("notificationRepository").(repository.NotificationRepository)}
}

type DeleteNotificationService struct {
	repository repository.NotificationRepository
}

func (u *DeleteNotificationService) DeleteNotification(notificationID int) error {
	return u.repository.DeleteNotification(notificationID)
}
