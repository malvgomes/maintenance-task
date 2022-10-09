package repository

import (
	"context"
	"maintenance-task/pkg/task/model"
)

type TaskRepository interface {
	CreateTask(input model.CreateTask) error
	UpdateTask(input model.UpdateTask) error
	DeleteTask(taskID int) error
	ListTasks(userID int) ([]*model.Task, error)
}

func GetTaskRepository(ctx context.Context) TaskRepository {
	return NewTaskRepositoryMariaDB(ctx)
}
