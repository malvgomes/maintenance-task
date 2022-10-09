package viewmodel

import (
	"maintenance-task/pkg/task/model"
	"time"
)

type Task struct {
	ID        int        `json:"id"`
	UserID    int        `json:"userId"`
	Summary   string     `json:"summary"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func mapToTaskViewmodel(t *model.Task) *Task {
	return &Task{
		ID:        t.ID,
		UserID:    t.UserID,
		Summary:   t.Summary,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func MapToTaskListViewmodel(t []*model.Task) []*Task {
	var tasks = make([]*Task, 0)

	for _, task := range t {
		tasks = append(tasks, mapToTaskViewmodel(task))
	}

	return tasks
}
