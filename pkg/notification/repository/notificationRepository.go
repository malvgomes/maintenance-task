package repository

import (
	"context"
	"maintenance-task/pkg/notification/model"
)

type NotificationRepository interface {
	CreateNotification(input model.CreateNotification) error
	DeleteNotification(notificationID int) error
	ClearNotifications(userID int) error
	ListNotifications(userID int) ([]*model.Notification, error)
}

func GetNotificationRepository(ctx context.Context) NotificationRepository {
	return NewNotificationRepositoryMariaDB(ctx)
}
