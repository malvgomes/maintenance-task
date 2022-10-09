package repository_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/notification/model"
	"maintenance-task/pkg/notification/repository"
	databaseMock "maintenance-task/shared/mock/database"
	"regexp"
	"testing"
	"time"

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

		err := repoMock.CreateNotification(input)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123, 456).
			WillReturnError(errors.New("SQL Failure on INSERT notification"))

		err := repoMock.CreateNotification(input)

		assert.EqualError(t, err, "SQL Failure on INSERT notification")
	})
}

func TestNotificationRepositoryMariaDB_DeleteNotification(t *testing.T) {
	query := regexp.QuoteMeta(`
		DELETE FROM maintenance.notifications WHERE id = ?;
	`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.DeleteNotification(123)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnError(errors.New("SQL Failure on DELETE notification by ID"))

		err := repoMock.DeleteNotification(123)

		assert.EqualError(t, err, "SQL Failure on DELETE notification by ID")
	})
}

func TestNotificationRepositoryMariaDB_ClearNotifications(t *testing.T) {
	query := regexp.QuoteMeta(`
		DELETE FROM maintenance.notifications WHERE user_id = ?;
	`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.ClearNotifications(123)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnError(errors.New("SQL Failure on DELETE notifications"))

		err := repoMock.ClearNotifications(123)

		assert.EqualError(t, err, "SQL Failure on DELETE notifications")
	})
}

func TestNotificationRepositoryMariaDB_ListNotifications(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT
    		id AS ID,
    		user_id AS UserID,
    		task_id AS TaskID,
    		is_update AS IsUpdate,
    		created_at AS CreatedAt
		FROM maintenance.notifications WHERE user_id = ?;
	`)

	loc, err := time.LoadLocation("America/Sao_Paulo")
	assert.NoError(t, err)

	date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectQuery(query).
			WithArgs(123).
			WillReturnRows(sqlmock.NewRows(
				[]string{"ID", "UserID", "TaskID", "IsUpdate", "CreatedAt"},
			).AddRow(444, 123, 321, 0, date))

		result, err := repoMock.ListNotifications(123)

		assert.NoError(t, err)
		assert.Equal(t, []*model.Notification{{
			ID:        444,
			UserID:    123,
			TaskID:    321,
			IsUpdate:  false,
			CreatedAt: date,
		}}, result)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedNotificationRepository(t)

		dbMock.ExpectQuery(query).
			WithArgs(123).
			WillReturnError(errors.New("SQL Failure on SELECT notifications"))

		result, err := repoMock.ListNotifications(123)

		assert.EqualError(t, err, "SQL Failure on SELECT notifications")
		assert.Empty(t, result)
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
