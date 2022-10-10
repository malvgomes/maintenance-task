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

func (n *notificationRepositoryMariaDB) CreateNotification(input model.CreateNotification) (int, error) {
	res, err := n.db.Exec(n.queries["insert-notification"], input.UserID, input.TaskID)
	if err != nil {
		return 0, err
	}

	ID, _ := res.LastInsertId()

	return int(ID), nil
}

func NewNotificationRepositoryMariaDB(ctx context.Context) NotificationRepository {
	return &notificationRepositoryMariaDB{
		queries: goyesql.MustParseBytes(queries),
		db:      ctx.Value("database").(database.Database),
	}
}
