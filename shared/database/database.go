package database

import (
	"database/sql"
)

type Database interface {
	SelectOne(interface{}, string, ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
}

func GetDatabase() (Database, error) {
	return Open()
}
