package repository_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/repository"
	databaseMock "maintenance-task/shared/mock/database"
	"maintenance-task/shared/pointer"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryMariaDB_CreateUser(t *testing.T) {
	query := regexp.QuoteMeta(`
		INSERT INTO maintenance.users (username, password, user_first_name, user_last_name, user_role)
		VALUES (?, AES_ENCRYPT(?, 'secure_key'), ?, ?, ?);
	`)

	input := model.CreateUser{
		Username:  "userName",
		Password:  "password",
		FirstName: "firstName",
		LastName:  pointer.String("lastName"),
		UserRole:  "MANAGER",
	}

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectExec(query).
			WithArgs("userName", "password", "firstName", "lastName", "MANAGER").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.CreateUser(input)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectExec(query).
			WithArgs("userName", "password", "firstName", "lastName", "MANAGER").
			WillReturnError(errors.New("SQL Failure on INSERT"))

		err := repoMock.CreateUser(input)

		assert.EqualError(t, err, "SQL Failure on INSERT")
	})
}

func TestUserRepositoryMariaDB_DeleteUser(t *testing.T) {
	query := regexp.QuoteMeta(`DELETE FROM maintenance.users WHERE username = ?;`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectExec(query).
			WithArgs("userName").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.DeleteUser("userName")

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectExec(query).
			WithArgs("userName").
			WillReturnError(errors.New("SQL failure on DELETE"))

		err := repoMock.DeleteUser("userName")

		assert.EqualError(t, err, "SQL failure on DELETE")
	})
}

func getMockedUserRepository(t *testing.T) (repository.UserRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := databaseMock.NewDatabaseMock()
	assert.NoError(t, err)

	ctx := context.WithValue(context.Background(), "database", db)

	return repository.NewUserRepositoryMariaDB(ctx), mock
}
