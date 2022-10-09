package service

import (
	"context"
	"maintenance-task/pkg/notification/model"
	"maintenance-task/pkg/notification/repository"
)

func NewListNotificationsService(ctx context.Context) *ListNotificationsService {
	return &ListNotificationsService{
		repository: ctx.Value("notificationRepository").(repository.NotificationRepository)}
}

type ListNotificationsService struct {
	repository repository.NotificationRepository
}

func (u *ListNotificationsService) ListNotifications(userID int) ([]*model.Notification, error) {
	return u.repository.ListNotifications(userID)
}
