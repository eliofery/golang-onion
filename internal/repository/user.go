package repository

import (
	"database/sql"
	"github.com/eliofery/golang-angular/internal/dto"
)

// UserQuery содержит запросы в базу данных для манипуляции с пользователями
type UserQuery interface {
	CreateUser(user dto.UserCreate) (*int, error)
}

type userQuery struct {
	db *sql.DB
}

func (u *userQuery) CreateUser(user dto.UserCreate) (*int, error) {
	query := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int
	err := u.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
