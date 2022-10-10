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

		repoMock.EXPECT().CreateUser(input).Return(1, nil)

		ID, err := svc.CreateUser(input)

		assert.NoError(t, err)
		assert.Equal(t, 1, ID)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateUserService(t)
		defer finish()

		repoMock.EXPECT().CreateUser(input).Return(0, errors.New("repository error"))

		ID, err := svc.CreateUser(input)

		assert.EqualError(t, err, "repository error")
		assert.Equal(t, 0, ID)
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
