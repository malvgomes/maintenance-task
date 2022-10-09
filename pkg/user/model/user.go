package model

import "time"

type User struct {
	Username  string
	Password  string
	FirstName string
	LastName  *string
	UserRole  Role
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type CreateUser struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	FirstName string  `json:"firstName"`
	LastName  *string `json:"lastName"`
	UserRole  Role    `json:"userRole"`
}

func (u *CreateUser) IsValid() bool {
	return u != nil && u.Username != "" && u.Password != "" && u.FirstName != "" && u.UserRole.
		IsValid()
}
