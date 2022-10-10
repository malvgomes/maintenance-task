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

func TestGetTaskService_GetTask(t *testing.T) {
	input := 123

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetTaskService(t)
		defer finish()

		repoMock.EXPECT().GetTask(input).Return(&model.Task{ID: 1, Summary: "Summary"}, nil)

		res, err := svc.GetTask(input)

		assert.NoError(t, err)
		assert.Equal(t, &model.Task{ID: 1, Summary: "Summary"}, res)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedGetTaskService(t)
		defer finish()

		repoMock.EXPECT().GetTask(input).Return(nil, errors.New("repository error"))

		res, err := svc.GetTask(input)

		assert.EqualError(t, err, "repository error")
		assert.Empty(t, res)
	})
}

func getMockedGetTaskService(t *testing.T) (
	*service.GetTaskService,
	*mockRepository.MockTaskRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockTaskRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewGetTaskService(
			context.WithValue(context.Background(), "taskRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
