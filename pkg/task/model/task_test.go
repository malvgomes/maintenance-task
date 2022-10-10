package model_test

import (
	"maintenance-task/pkg/task/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		input *model.CreateTask
		valid bool
	}{
		{"valid", &model.CreateTask{UserID: 123, Summary: "3"}, true},
		{"nil", nil, false},
		{"empty", &model.CreateTask{}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.valid, test.input.IsValid())
		})
	}
}

func TestUpdateTask_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		input *model.UpdateTask
		valid bool
	}{
		{"valid", &model.UpdateTask{ID: 123, UserID: 123, Summary: "3"}, true},
		{"nil", nil, false},
		{"empty", &model.UpdateTask{}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.valid, test.input.IsValid())
		})
	}
}
