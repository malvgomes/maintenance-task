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
	"time"

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

		ID, err := repoMock.CreateUser(input)

		assert.NoError(t, err)
		assert.Equal(t, 1, ID)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectExec(query).
			WithArgs("userName", "password", "firstName", "lastName", "MANAGER").
			WillReturnError(errors.New("SQL Failure on INSERT"))

		ID, err := repoMock.CreateUser(input)

		assert.EqualError(t, err, "SQL Failure on INSERT")
		assert.Equal(t, 0, ID)
	})
}

func TestUserRepositoryMariaDB_DeleteUser(t *testing.T) {
	query := regexp.QuoteMeta(`DELETE FROM maintenance.users WHERE id = ?;`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.DeleteUser(123)

		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectExec(query).
			WithArgs(123).
			WillReturnError(errors.New("SQL failure on DELETE"))

		err := repoMock.DeleteUser(123)

		assert.EqualError(t, err, "SQL failure on DELETE")
	})
}

func TestUserRepositoryMariaDB_GetUser(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT
    		id AS ID,
    		username AS Username,
    		user_first_name AS FirstName,
    		user_last_name AS LastName,
    		user_role AS UserRole,
    		created_at AS CreatedAt,
    		updated_at AS UpdatedAt
		FROM maintenance.users WHERE username = ? AND password = AES_ENCRYPT(?, 'secure_key');
	`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.NoError(t, err)

		date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

		dbMock.ExpectQuery(query).
			WithArgs("user", "pass").
			WillReturnRows(sqlmock.NewRows([]string{
				"ID", "Username", "FirstName", "LastName", "UserRole", "CreatedAt", "UpdatedAt"}).
				AddRow(1, "user", "first", "last", "MANAGER", date, nil))

		user, err := repoMock.GetUser("user", "pass")

		assert.NoError(t, err)
		assert.Equal(t, &model.User{
			ID:        1,
			Username:  "user",
			FirstName: "first",
			LastName:  pointer.String("last"),
			UserRole:  model.Manager,
			CreatedAt: date,
			UpdatedAt: nil,
		}, user)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectQuery(query).
			WithArgs("user", "pass").
			WillReturnError(errors.New("SQL failure on SELECT user"))

		user, err := repoMock.GetUser("user", "pass")

		assert.EqualError(t, err, "SQL failure on SELECT user")
		assert.Empty(t, user)
	})
}

func TestUserRepositoryMariaDB_GetUserByID(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT
		    id AS ID,
		    username AS Username,
		    user_first_name AS FirstName,
		    user_last_name AS LastName,
		    user_role AS UserRole,
		    created_at AS CreatedAt,
		    updated_at AS UpdatedAt
		FROM maintenance.users WHERE id = ?;
	`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.NoError(t, err)

		date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

		dbMock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{
				"ID", "Username", "FirstName", "LastName", "UserRole", "CreatedAt", "UpdatedAt"}).
				AddRow(1, "user", "first", "last", "MANAGER", date, nil))

		user, err := repoMock.GetUserByID(1)

		assert.NoError(t, err)
		assert.Equal(t, &model.User{
			ID:        1,
			Username:  "user",
			FirstName: "first",
			LastName:  pointer.String("last"),
			UserRole:  model.Manager,
			CreatedAt: date,
			UpdatedAt: nil,
		}, user)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectQuery(query).
			WithArgs(1).
			WillReturnError(errors.New("SQL failure on SELECT user by ID"))

		user, err := repoMock.GetUserByID(1)

		assert.EqualError(t, err, "SQL failure on SELECT user by ID")
		assert.Empty(t, user)
	})
}

func TestUserRepositoryMariaDB_GetUsersByRole(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT
		    id AS ID,
		    username AS Username,
		    user_first_name AS FirstName,
		    user_last_name AS LastName,
		    user_role AS UserRole,
		    created_at AS CreatedAt,
		    updated_at AS UpdatedAt
		FROM maintenance.users WHERE user_role = ?;
	`)

	t.Run("Success", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.NoError(t, err)

		date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

		dbMock.ExpectQuery(query).
			WithArgs(model.Manager).
			WillReturnRows(sqlmock.NewRows([]string{
				"ID", "Username", "FirstName", "LastName", "UserRole", "CreatedAt", "UpdatedAt"}).
				AddRow(1, "user", "first", "last", "MANAGER", date, nil))

		user, err := repoMock.GetUsersByRole(model.Manager)

		assert.NoError(t, err)
		assert.Equal(t, []*model.User{{
			ID:        1,
			Username:  "user",
			FirstName: "first",
			LastName:  pointer.String("last"),
			UserRole:  model.Manager,
			CreatedAt: date,
			UpdatedAt: nil,
		}}, user)
	})

	t.Run("Failure", func(t *testing.T) {
		repoMock, dbMock := getMockedUserRepository(t)

		dbMock.ExpectQuery(query).
			WithArgs(model.Manager).
			WillReturnError(errors.New("SQL failure on SELECT users by role"))

		user, err := repoMock.GetUsersByRole(model.Manager)

		assert.EqualError(t, err, "SQL failure on SELECT users by role")
		assert.Empty(t, user)
	})
}

func getMockedUserRepository(t *testing.T) (repository.UserRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := databaseMock.NewDatabaseMock()
	assert.NoError(t, err)

	ctx := context.WithValue(context.Background(), "database", db)

	return repository.NewUserRepositoryMariaDB(ctx), mock
}
