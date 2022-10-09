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

func TestCreateTaskService_CreateTask(t *testing.T) {
	input := model.CreateTask{
		UserID:  123,
		Summary: "summary",
	}

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateTaskService(t)
		defer finish()

		repoMock.EXPECT().CreateTask(input).Return(nil)

		assert.NoError(t, svc.CreateTask(input))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateTaskService(t)
		defer finish()

		repoMock.EXPECT().CreateTask(input).Return(errors.New("repository error"))

		assert.EqualError(t, svc.CreateTask(input), "repository error")
	})
}

func getMockedCreateTaskService(t *testing.T) (
	*service.CreateTaskService,
	*mockRepository.MockTaskRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockTaskRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewCreateTaskService(
			context.WithValue(context.Background(), "taskRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
