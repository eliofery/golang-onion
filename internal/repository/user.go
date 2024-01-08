package repository

import (
	"database/sql"
	"errors"
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// UserQuery содержит запросы в базу данных для манипуляции с пользователями
type UserQuery interface {
	Save(user dto.UserCreate) (id int, err error)
}

type userQuery struct {
	db *sql.DB
}

// Save создание пользователя
func (u *userQuery) Save(user dto.UserCreate) (id int, err error) {
	query := "INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"

	err = u.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, errors.New("пользователь уже существует")
		}
		return 0, err
	}

	return id, nil
}
