package repository_test

import (
	"context"
	"maintenance-task/pkg/notification/repository"
	databaseMock "maintenance-task/shared/mock/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNotificationRepository(t *testing.T) {
	db, _, err := databaseMock.NewDatabaseMock()
	assert.NoError(t, err)

	assert.NotEmpty(
		t,
		repository.GetNotificationRepository(context.WithValue(context.Background(), "database", db)),
	)
}
