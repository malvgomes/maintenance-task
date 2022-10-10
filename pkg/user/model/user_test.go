package model_test

import (
	"maintenance-task/pkg/user/model"
	"maintenance-task/shared/pointer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		input *model.CreateUser
		valid bool
	}{
		{
			"Valid - All values", &model.CreateUser{
				Username:  "username",
				Password:  "pass",
				FirstName: "firstName",
				LastName:  pointer.String("lastName"),
				UserRole:  model.Technician,
			},
			true,
		},
		{
			"Valid - required values", &model.CreateUser{
				Username:  "username",
				Password:  "pass",
				FirstName: "firstName",
				UserRole:  model.Technician,
			},
			true,
		},
		{
			"Invalid - received nil",
			nil,
			false,
		},
		{
			"Invalid - received empty",
			&model.CreateUser{},
			false,
		},
		{
			"Invalid - received invalid role",
			&model.CreateUser{
				Username:  "username",
				Password:  "pass",
				FirstName: "firstName",
				UserRole:  "invalid",
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.valid, test.input.IsValid())
		})
	}
}
