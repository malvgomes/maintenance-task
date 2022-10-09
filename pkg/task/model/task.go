package model

import "time"

type Task struct {
	ID        int
	UserID    int
	Summary   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type CreateTask struct {
	UserID  int    `json:"userId"`
	Summary string `json:"summary"`
}

type UpdateTask struct {
	ID      int    `json:"id"`
	Summary string `json:"summary"`
}
