package service_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/service"
	mockRepository "maintenance-task/shared/mock/task/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestUpdateTaskService_UpdateTask(t *testing.T) {
	input := model.UpdateTask{
		ID:      123,
		Summary: "summary",
	}

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedUpdateTaskService(t)
		defer finish()

		repoMock.EXPECT().UpdateTask(input).Return(nil)

		assert.NoError(t, svc.UpdateTask(input))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedUpdateTaskService(t)
		defer finish()

		repoMock.EXPECT().UpdateTask(input).Return(errors.New("repository error"))

		assert.EqualError(t, svc.UpdateTask(input), "repository error")
	})
}

func getMockedUpdateTaskService(t *testing.T) (
	*service.UpdateTaskService,
	*mockRepository.MockTaskRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockTaskRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewUpdateTaskService(
			context.WithValue(context.Background(), "taskRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
