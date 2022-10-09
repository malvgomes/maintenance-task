package service_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/user/service"
	mock_repository "maintenance-task/shared/mock/user/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestDeleteUserService_DeleteUser(t *testing.T) {
	input := 123

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedDeleteUserService(t)
		defer finish()

		repoMock.EXPECT().DeleteUser(input).Return(nil)

		assert.NoError(t, svc.DeleteUser(input))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedDeleteUserService(t)
		defer finish()

		repoMock.EXPECT().DeleteUser(input).Return(errors.New("repository error"))

		assert.EqualError(t, svc.DeleteUser(input), "repository error")
	})
}

func getMockedDeleteUserService(t *testing.T) (
	*service.DeleteUserService,
	*mock_repository.MockUserRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mock_repository.NewMockUserRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewDeleteUserService(
			context.WithValue(context.Background(), "userRepository", repositoryMock),
		),
		repositoryMock,
		finish

}
