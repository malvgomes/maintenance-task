package repository_test

import (
	"context"
	"maintenance-task/pkg/task/repository"
	databaseMock "maintenance-task/shared/mock/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaskRepository(t *testing.T) {
	db, _, err := databaseMock.NewDatabaseMock()
	assert.NoError(t, err)

	assert.NotEmpty(
		t,
		repository.GetTaskRepository(context.WithValue(context.Background(), "database", db)),
	)
}
