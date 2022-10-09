package model

import "time"

type Notification struct {
	ID        int
	UserID    int
	TaskID    int
	IsUpdate  bool
	CreatedAt time.Time
}

type CreateNotification struct {
	UserID int
	TaskID int
}
