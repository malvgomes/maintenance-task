package model_test

import (
	"maintenance-task/pkg/user/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRole_IsValid(t *testing.T) {
	tests := []struct {
		name  string
		role  model.Role
		valid bool
	}{
		{"manager", model.Manager, true},
		{"technician", model.Technician, true},
		{"invalid", "Invalid", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.valid, test.role.IsValid())
		})
	}
}
