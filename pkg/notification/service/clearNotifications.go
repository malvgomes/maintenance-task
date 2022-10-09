package service

import (
	"context"
	"maintenance-task/pkg/notification/repository"
)

func NewClearNotificationsService(ctx context.Context) *ClearNotificationsService {
	return &ClearNotificationsService{
		repository: ctx.Value("notificationRepository").(repository.NotificationRepository)}
}

type ClearNotificationsService struct {
	repository repository.NotificationRepository
}

func (u *ClearNotificationsService) ClearNotifications(userID int) error {
	return u.repository.ClearNotifications(userID)
}
