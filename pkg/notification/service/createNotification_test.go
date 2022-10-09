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

		repoMock.EXPECT().CreateNotification(input).Return(nil)

		assert.NoError(t, svc.CreateNotification(input))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedCreateNotificationService(t)
		defer finish()

		repoMock.EXPECT().CreateNotification(input).Return(errors.New("repository error"))

		assert.EqualError(t, svc.CreateNotification(input), "repository error")
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
