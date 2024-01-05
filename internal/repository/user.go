package repository

import (
	"database/sql"
)

// UserQuery содержит запросы в базу данных для манипуляции с пользователями
type UserQuery interface {
	GetUser() error
}

type userQuery struct {
	db *sql.DB
}

func (u *userQuery) GetUser() error {
	query := "SELECT ..."
	_ = query

	return nil
}
