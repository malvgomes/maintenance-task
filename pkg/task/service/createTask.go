package service

import (
	"context"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/repository"
)

func NewCreateTaskService(ctx context.Context) *CreateTaskService {
	return &CreateTaskService{repository: ctx.Value("taskRepository").(repository.TaskRepository)}
}

type CreateTaskService struct {
	repository repository.TaskRepository
}

func (t *CreateTaskService) CreateTask(input model.CreateTask) (int, error) {
	return t.repository.CreateTask(input)
}
