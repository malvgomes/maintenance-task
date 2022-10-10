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

		repoMock.EXPECT().CreateTask(input).Return(1, nil)

		ID, err := svc.CreateTask(input)

		assert.NoError(t, err)
		assert.Equal(t, 1, ID)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateTaskService(t)
		defer finish()

		repoMock.EXPECT().CreateTask(input).Return(0, errors.New("repository error"))

		ID, err := svc.CreateTask(input)

		assert.EqualError(t, err, "repository error")
		assert.Equal(t, 0, ID)
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
