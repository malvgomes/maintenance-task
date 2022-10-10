package service

import (
	"context"
	"maintenance-task/pkg/task/model"
	"maintenance-task/pkg/task/repository"
)

func NewGetTaskService(ctx context.Context) *GetTaskService {
	return &GetTaskService{repository: ctx.Value("taskRepository").(repository.TaskRepository)}
}

type GetTaskService struct {
	repository repository.TaskRepository
}

func (t *GetTaskService) GetTask(userID int) (*model.Task, error) {
	return t.repository.GetTask(userID)
}
