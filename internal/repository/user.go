package repository

import (
	"database/sql"
	"github.com/eliofery/golang-angular/internal/model"
)

// UserQuery содержит запросы в базу данных для манипуляции с пользователями
type UserQuery interface {
	CreateUser(user model.User) (*int, error)
}

type userQuery struct {
	db *sql.DB
}

func (u *userQuery) CreateUser(user model.User) (*int, error) {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	var id int
	err := u.db.QueryRow(query, user.Email, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
