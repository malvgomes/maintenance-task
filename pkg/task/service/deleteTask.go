package service

import (
	"context"
	"maintenance-task/pkg/task/repository"
)

func NewDeleteTaskService(ctx context.Context) *DeleteTaskService {
	return &DeleteTaskService{repository: ctx.Value("taskRepository").(repository.TaskRepository)}
}

type DeleteTaskService struct {
	repository repository.TaskRepository
}

func (t *DeleteTaskService) DeleteTask(taskID int) error {
	return t.repository.DeleteTask(taskID)
}
