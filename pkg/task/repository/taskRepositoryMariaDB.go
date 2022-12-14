package repository

import (
	"context"
	_ "embed"
	"maintenance-task/pkg/task/model"
	"maintenance-task/shared/database"

	"github.com/nleof/goyesql"
)

//go:embed queries.sql
var queries []byte

type taskRepositoryMariaDB struct {
	queries goyesql.Queries
	db      database.Database
}

func (t *taskRepositoryMariaDB) CreateTask(input model.CreateTask) (int, error) {
	res, err := t.db.Exec(t.queries["insert-task"], input.UserID, input.Summary)
	if err != nil {
		return 0, err
	}

	ID, _ := res.LastInsertId()

	return int(ID), nil
}

func (t *taskRepositoryMariaDB) UpdateTask(input model.UpdateTask) error {
	_, err := t.db.Exec(t.queries["update-task"], input.Summary, input.ID)

	return err
}

func (t *taskRepositoryMariaDB) DeleteTask(taskID int) error {
	_, err := t.db.Exec(t.queries["delete-task"], taskID)

	return err
}

func (t *taskRepositoryMariaDB) ListTasks(userID int) ([]*model.Task, error) {
	var tasks []*model.Task

	err := t.db.Select(&tasks, t.queries["list-tasks"], userID)

	return tasks, err
}

func (t *taskRepositoryMariaDB) GetTask(taskID int) (*model.Task, error) {
	var task *model.Task

	err := t.db.SelectOne(&task, t.queries["get-task"], taskID)

	return task, err
}

func NewTaskRepositoryMariaDB(ctx context.Context) TaskRepository {
	return &taskRepositoryMariaDB{
		queries: goyesql.MustParseBytes(queries),
		db:      ctx.Value("database").(database.Database),
	}
}
