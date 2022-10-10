package service_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/notification/model"
	"maintenance-task/pkg/notification/service"
	mockRepository "maintenance-task/shared/mock/notification/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestCreateNotificationService_CreateNotification(t *testing.T) {
	input := model.CreateNotification{
		UserID: 123,
		TaskID: 456,
	}

	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateNotificationService(t)
		defer finish()

		repoMock.EXPECT().CreateNotification(input).Return(1, nil)

		ID, err := svc.CreateNotification(input)

		assert.Equal(t, 1, ID)
		assert.NoError(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateNotificationService(t)
		defer finish()

		repoMock.EXPECT().CreateNotification(input).Return(0, errors.New("repository error"))

		ID, err := svc.CreateNotification(input)

		assert.Equal(t, 0, ID)
		assert.EqualError(t, err, "repository error")
	})
}

func getMockedCreateNotificationService(t *testing.T) (
	*service.CreateNotificationService,
	*mockRepository.MockNotificationRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockNotificationRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewCreateNotificationService(
			context.WithValue(context.Background(), "notificationRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
