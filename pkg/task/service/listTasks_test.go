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

func TestListTasksService_ListTasks(t *testing.T) {
	input := 123

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedListTasksService(t)
		defer finish()

		repoMock.EXPECT().ListTasks(input).Return([]*model.Task{{ID: 1}, {ID: 2}}, nil)

		res, err := svc.ListTasks(input)

		assert.NoError(t, err)
		assert.Len(t, res, 2)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedListTasksService(t)
		defer finish()

		repoMock.EXPECT().ListTasks(input).Return(nil, errors.New("repository error"))

		res, err := svc.ListTasks(input)

		assert.EqualError(t, err, "repository error")
		assert.Empty(t, res)
	})
}

func getMockedListTasksService(t *testing.T) (
	*service.ListTasksService,
	*mockRepository.MockTaskRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockTaskRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewListTasksService(
			context.WithValue(context.Background(), "taskRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
