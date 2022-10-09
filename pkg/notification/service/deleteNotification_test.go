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

func TestDeleteNotificationService_DeleteNotification(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedDeleteNotificationService(t)
		defer finish()

		repoMock.EXPECT().DeleteNotification(12345).Return(nil)

		assert.NoError(t, svc.DeleteNotification(12345))
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedDeleteNotificationService(t)
		defer finish()

		repoMock.EXPECT().DeleteNotification(12345).Return(errors.New("repository error"))

		assert.EqualError(t, svc.DeleteNotification(12345), "repository error")
	})
}

func getMockedDeleteNotificationService(t *testing.T) (
	*service.DeleteNotificationService,
	*mockRepository.MockNotificationRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockNotificationRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewDeleteNotificationService(
			context.WithValue(context.Background(), "notificationRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
