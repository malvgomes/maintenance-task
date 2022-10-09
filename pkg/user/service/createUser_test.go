package service_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/user/model"
	"maintenance-task/pkg/user/service"
	mock_repository "maintenance-task/shared/mock/user/repository"
	"maintenance-task/shared/pointer"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestCreateUserService_CreateUser(t *testing.T) {
	input := model.CreateUser{
		Username:  "userName",
		Password:  "password",
		FirstName: "firstName",
		LastName:  pointer.String("lastName"),
		UserRole:  "MANAGER",
	}

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateUserService(t)
		defer finish()

		repoMock.EXPECT().CreateUser(input).Return(nil)

		assert.NoError(t, svc.CreateUser(input))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateUserService(t)
		defer finish()

		repoMock.EXPECT().CreateUser(input).Return(errors.New("repository error"))

		assert.EqualError(t, svc.CreateUser(input), "repository error")
	})
}

func getMockedCreateUserService(t *testing.T) (
	*service.CreateUserService,
	*mock_repository.MockUserRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mock_repository.NewMockUserRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewCreateUserService(
			context.WithValue(context.Background(), "userRepository", repositoryMock),
		),
		repositoryMock,
		finish

}
