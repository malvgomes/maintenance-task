package database

import (
	"database/sql"
)

type Database interface {
	Select(i interface{}, s string, args ...interface{}) error
	SelectOne(interface{}, string, ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
}

func GetDatabase() (Database, error) {
	return Open()
}
