package repository

import (
	"context"
	_ "embed"
	"maintenance-task/pkg/notification/model"
	"maintenance-task/shared/database"

	"github.com/nleof/goyesql"
)

//go:embed queries.sql
var queries []byte

type notificationRepositoryMariaDB struct {
	queries goyesql.Queries
	db      database.Database
}

func (n notificationRepositoryMariaDB) CreateNotification(input model.CreateNotification) error {
	_, err := n.db.Exec(n.queries["insert-notification"], input.UserID, input.TaskID)
	return err
}

func (n notificationRepositoryMariaDB) DeleteNotification(notificationID int) error {
	_, err := n.db.Exec(n.queries["delete-notification"], notificationID)
	return err
}

func (n notificationRepositoryMariaDB) ClearNotifications(userID int) error {
	_, err := n.db.Exec(n.queries["clear-notifications"], userID)
	return err
}

func (n notificationRepositoryMariaDB) ListNotifications(userID int) ([]*model.Notification, error) {
	var notifications []*model.Notification

	err := n.db.Select(&notifications, n.queries["list-notifications"], userID)

	return notifications, err
}

func NewNotificationRepositoryMariaDB(ctx context.Context) NotificationRepository {
	return &notificationRepositoryMariaDB{
		queries: goyesql.MustParseBytes(queries),
		db:      ctx.Value("database").(database.Database),
	}
}
