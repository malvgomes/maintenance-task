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

func TestListNotificationsService_ListNotifications(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		svc, repoMock, finish := getMockedListNotificationsService(t)
		defer finish()

		repoMock.EXPECT().ListNotifications(123).Return([]*model.Notification{
			{ID: 1}, {ID: 2},
		}, nil)

		res, err := svc.ListNotifications(123)
		assert.NoError(t, err)
		assert.Len(t, res, 2)
	})

	t.Run("Failure", func(t *testing.T) {
		svc, repoMock, finish := getMockedListNotificationsService(t)
		defer finish()

		repoMock.EXPECT().ListNotifications(123).Return(nil, errors.New("repository error"))

		res, err := svc.ListNotifications(123)
		assert.EqualError(t, err, "repository error")
		assert.Empty(t, res)
	})
}

func getMockedListNotificationsService(t *testing.T) (
	*service.ListNotificationsService,
	*mockRepository.MockNotificationRepository,
	func(),
) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := mockRepository.NewMockNotificationRepository(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return service.NewListNotificationsService(
			context.WithValue(context.Background(), "notificationRepository", repositoryMock),
		),
		repositoryMock,
		finish
}
