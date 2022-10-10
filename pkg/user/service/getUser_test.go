package service_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/service"
	mock_repository "maintenance-task/shared/mock/user/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestGetUserService_GetUser(t *testing.T) {
	username := "user"
	password := "pass"

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetUserService(t)
		defer finish()

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.NoError(t, err)

		date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

		repoMock.EXPECT().GetUser(username, password).Return(&model.User{
			ID:        123,
			Username:  "user",
			FirstName: "User",
			UserRole:  model.Technician,
			CreatedAt: date,
			UpdatedAt: nil,
		}, nil)

		user, err := svc.GetUser(username, password)

		assert.NoError(t, err)
		assert.Equal(t, &model.User{
			ID:        123,
			Username:  "user",
			FirstName: "User",
			UserRole:  model.Technician,
			CreatedAt: date,
			UpdatedAt: nil,
		}, user)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetUserService(t)
		defer finish()

		repoMock.EXPECT().GetUser(username, password).Return(nil, errors.New("repository error"))

		user, err := svc.GetUser(username, password)

		assert.EqualError(t, err, "repository error")
		assert.Empty(t, user)
	})
}

func TestGetUserService_GetUserByID(t *testing.T) {
	ID := 123

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetUserService(t)
		defer finish()

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.NoError(t, err)

		date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

		repoMock.EXPECT().GetUserByID(ID).Return(&model.User{
			ID:        123,
			Username:  "user",
			FirstName: "User",
			UserRole:  model.Technician,
			CreatedAt: date,
			UpdatedAt: nil,
		}, nil)

		user, err := svc.GetUserByID(ID)

		assert.NoError(t, err)
		assert.Equal(t, &model.User{
			ID:        123,
			Username:  "user",
			FirstName: "User",
			UserRole:  model.Technician,
			CreatedAt: date,
			UpdatedAt: nil,
		}, user)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetUserService(t)
		defer finish()

		repoMock.EXPECT().GetUserByID(ID).Return(nil, errors.New("repository error"))

		user, err := svc.GetUserByID(ID)

		assert.EqualError(t, err, "repository error")
		assert.Empty(t, user)
	})
}

func TestGetUserService_GetUsersByRole(t *testing.T) {
	role := model.Technician

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetUserService(t)
		defer finish()

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.NoError(t, err)

		date := time.Date(2022, 9, 9, 11, 12, 13, 0, loc)

		repoMock.EXPECT().GetUsersByRole(role).Return([]*model.User{{
			ID:        123,
			Username:  "user",
			FirstName: "User",
			UserRole:  role,
			CreatedAt: date,
			UpdatedAt: nil,
		}}, nil)

		user, err := svc.GetUsersByRole(role)

		assert.NoError(t, err)
		assert.Equal(t, []*model.User{{
			ID:        123,
			Username:  "user",
			FirstName: "User",
			UserRole:  role,
			CreatedAt: date,
			UpdatedAt: nil,
		}}, user)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetUserService(t)
		defer finish()

		repoMock.EXPECT().GetUsersByRole(role).Return(nil, errors.New("repository error"))

		user, err := svc.GetUsersByRole(role)

		assert.EqualError(t, err, "repository error")
		assert.Empty(t, user)
	})
}

func getMockedGetUserService(t *testing.T) (
	*service.GetUserService,
	*mock_repository.MockUserRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mock_repository.NewMockUserRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewGetUserService(
			context.WithValue(context.Background(), "userRepository", repositoryMock),
		),
		repositoryMock,
		finish

}
