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

func (t *CreateTask) IsValid() bool {
	return t != nil && t.UserID != 0 && t.Summary != "" && len(t.Summary) < 2500
}

type UpdateTask struct {
	ID      int    `json:"id"`
	Summary string `json:"summary"`
}

func (t *UpdateTask) IsValid() bool {
	return t != nil && t.ID != 0 && t.Summary != "" && len(t.Summary) < 2500
}
