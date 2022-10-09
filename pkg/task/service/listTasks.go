package service

import (
	"context"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/repository"
)

func NewListTasksService(ctx context.Context) *ListTasksService {
	return &ListTasksService{repository: ctx.Value("taskRepository").(repository.TaskRepository)}
}

type ListTasksService struct {
	repository repository.TaskRepository
}

func (t *ListTasksService) ListTasks(userID int) ([]*model.Task, error) {
	return t.repository.ListTasks(userID)
}
