package service_test

import (
	"context"
	"errors"
	"maintenance-task/pkg/notification/service"
	mockRepository "maintenance-task/shared/mock/notification/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestClearNotificationsService_ClearNotifications(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedClearNotificationsService(t)
		defer finish()

		repoMock.EXPECT().ClearNotifications(12345).Return(nil)

		assert.NoError(t, svc.ClearNotifications(12345))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedClearNotificationsService(t)
		defer finish()

		repoMock.EXPECT().ClearNotifications(12345).Return(errors.New("repository error"))

		assert.EqualError(t, svc.ClearNotifications(12345), "repository error")
	})
}

func getMockedClearNotificationsService(t *testing.T) (
	*service.ClearNotificationsService,
	*mockRepository.MockNotificationRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockNotificationRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewClearNotificationsService(
			context.WithValue(context.Background(), "notificationRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
