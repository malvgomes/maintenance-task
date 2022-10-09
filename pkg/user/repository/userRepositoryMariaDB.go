package repository

import (
	"context"
	_ "embed"
	"maintenance-task/pkg/user/model"
	"maintenance-task/shared/database"

	"github.com/nleof/goyesql"
)

//go:embed queries.sql
var queries []byte

type userRepositoryMariaDB struct {
	queries goyesql.Queries
	db      database.Database
}

func NewUserRepositoryMariaDB(ctx context.Context) UserRepository {
	return &userRepositoryMariaDB{
		queries: goyesql.MustParseBytes(queries),
		db:      ctx.Value("database").(database.Database),
	}
}

func (u *userRepositoryMariaDB) CreateUser(data model.CreateUser) error {
	_, err := u.db.Exec(u.queries["insert-user"], data.Username, data.Password, data.FirstName,
		data.LastName, data.UserRole)

	return err
}
