package service_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/task/service"
	mockRepository "maintenance-task/shared/mock/task/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestDeleteTaskService_DeleteTask(t *testing.T) {
	input := 123

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedDeleteTaskService(t)
		defer finish()

		repoMock.EXPECT().DeleteTask(input).Return(nil)

		assert.NoError(t, svc.DeleteTask(input))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedDeleteTaskService(t)
		defer finish()

		repoMock.EXPECT().DeleteTask(input).Return(errors.New("repository error"))

		assert.EqualError(t, svc.DeleteTask(input), "repository error")
	})
}

func getMockedDeleteTaskService(t *testing.T) (
	*service.DeleteTaskService,
	*mockRepository.MockTaskRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockTaskRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewDeleteTaskService(
			context.WithValue(context.Background(), "taskRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
