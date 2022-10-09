package repository_test

import (
	"context"
	"maintenance-task/pkg/user/repository"
	databaseMock "maintenance-task/shared/mock/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserRepository(t *testing.T) {
	db, _, err := databaseMock.NewDatabaseMock()
	assert.NoError(t, err)

	assert.NotEmpty(
		t,
		repository.GetUserRepository(context.WithValue(context.Background(), "database", db)),
	)
}
