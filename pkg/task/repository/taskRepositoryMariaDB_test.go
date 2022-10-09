package repository_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/repository"
	databaseMock "maintenance-task/shared/mock/database"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTaskRepositoryMariaDB_CreateTask(t *testing.T) {
	query := regexp.QuoteMeta(`
		INSERT INTO maintenance.tasks (user_id, summary)
		VALUES (?, AES_ENCRYPT(?, 'secure_key'));
	`)

	input := model.CreateTask{
		UserID:  123,
		Summary: "Task summary",
	}

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123, "Task summary").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.CreateTask(input)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123, "Task summary").
			WillReturnError(errors.New("SQL Failure on INSERT task"))

		err := repoMock.CreateTask(input)

		assert.EqualError(t, err, "SQL Failure on INSERT task")
	})
}

func TestTaskRepositoryMariaDB_DeleteTask(t *testing.T) {
	query := regexp.QuoteMeta(`
		DELETE FROM maintenance.tasks WHERE id = ?;
	`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.DeleteTask(123)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnError(errors.New("SQL Failure on DELETE task by ID"))

		err := repoMock.DeleteTask(123)

		assert.EqualError(t, err, "SQL Failure on DELETE task by ID")
	})
}

func TestTaskRepositoryMariaDB_UpdateTask(t *testing.T) {
	query := regexp.QuoteMeta(`
		UPDATE maintenance.tasks SET
    		summary = ?,
    		updated_at = NOW()
		WHERE id = ?;
	`)

	input := model.UpdateTask{
		ID:      456,
		Summary: "Task summary updated",
	}

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectExec(query).
			WithArgs("Task summary updated", 456).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.UpdateTask(input)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectExec(query).
			WithArgs("Task summary updated", 456).
			WillReturnError(errors.New("SQL Failure on UPDATE task"))

		err := repoMock.UpdateTask(input)

		assert.EqualError(t, err, "SQL Failure on UPDATE task")
	})
}

func TestTaskRepositoryMariaDB_ListTasks(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT
    		id AS ID,
    		user_id AS UserID,
    		AES_DECRYPT(summary, 'secure_key') AS Summary,
    		created_at AS CreatedAt,
    		updated_at AS UpdatedAt
		FROM maintenance.tasks WHERE user_id = ?;		
	`)

	loc, err := time.LoadLocation("America/Sao_Paulo")
	assert.NoError(t, err)

	date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectQuery(query).
			WithArgs(4567).
			WillReturnRows(sqlmock.NewRows(
				[]string{"ID", "UserID", "Summary", "CreatedAt", "UpdatedAt"},
			).AddRow(1234, 4567, "Summary", date, nil))

		result, err := repoMock.ListTasks(4567)

		assert.NoError(t, err)
		assert.Equal(t, []*model.Task{{
			ID:        1234,
			UserID:    4567,
			Summary:   "Summary",
			CreatedAt: date,
			UpdatedAt: nil,
		}}, result)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedTaskRepository(t)

		dbMock.ExpectQuery(query).
			WithArgs(4567).
			WillReturnError(errors.New("SQL Failure on SELECT tasks"))

		result, err := repoMock.ListTasks(4567)

		assert.EqualError(t, err, "SQL Failure on SELECT tasks")
		assert.Empty(t, result)
	})
}

func getMockedTaskRepository(t *testing.T) (
	repository.TaskRepository,
	sqlmock.Sqlmock,
) {
	t.Helper()

	db, mock, err := databaseMock.NewDatabaseMock()
	assert.NoError(t, err)

	ctx := context.WithValue(context.Background(), "database", db)

	return repository.NewTaskRepositoryMariaDB(ctx), mock
}
