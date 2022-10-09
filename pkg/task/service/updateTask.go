package service

import (
	"context"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/repository"
)

func NewUpdateTaskService(ctx context.Context) *UpdateTaskService {
	return &UpdateTaskService{repository: ctx.Value("taskRepository").(repository.TaskRepository)}
}

type UpdateTaskService struct {
	repository repository.TaskRepository
}

func (t *UpdateTaskService) UpdateTask(input model.UpdateTask) error {
	return t.repository.UpdateTask(input)
}
