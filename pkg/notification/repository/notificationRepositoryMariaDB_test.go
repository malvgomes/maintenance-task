package repository_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/notification/model"
	"maintenance-task/pkg/notification/repository"
	databaseMock "maintenance-task/shared/mock/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNotificationRepositoryMariaDB_CreateNotification(t *testing.T) {
	query := regexp.QuoteMeta(`
		INSERT INTO maintenance.notifications (user_id, task_id)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE is_update = 1;
	`)

	input := model.CreateNotification{
		UserID: 123,
		TaskID: 456,
	}

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123, 456).
			WillReturnResult(sqlmock.NewResult(1, 1))

		ID, err := repoMock.CreateNotification(input)

		assert.NoError(t, err)
		assert.Equal(t, 1, ID)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123, 456).
			WillReturnError(errors.New("SQL Failure on INSERT notification"))

		ID, err := repoMock.CreateNotification(input)

		assert.EqualError(t, err, "SQL Failure on INSERT notification")
		assert.Equal(t, 0, ID)
	})
}

func getMockedNotificationRepository(t *testing.T) (
	repository.NotificationRepository,
	sqlmock.Sqlmock,
) {
	t.Helper()

	db, mock, err := databaseMock.NewDatabaseMock()
	assert.NoError(t, err)

	ctx := context.WithValue(context.Background(), "database", db)

	return repository.NewNotificationRepositoryMariaDB(ctx), mock
}
