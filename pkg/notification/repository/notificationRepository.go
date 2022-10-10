package repository

import (
	"context"
	"maintenance-task/pkg/notification/model"
)

type NotificationRepository interface {
	CreateNotification(input model.CreateNotification) (int, error)
}

func GetNotificationRepository(ctx context.Context) NotificationRepository {
	return NewNotificationRepositoryMariaDB(ctx)
}
