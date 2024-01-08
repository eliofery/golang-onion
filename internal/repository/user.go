package repository

import (
	"database/sql"
	"errors"
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// UserQuery содержит запросы в базу данных для манипуляции с пользователями
type UserQuery interface {
	Save(user dto.UserCreate) (id int, err error)
	GetUserByEmail(email string) (user model.User, err error)
}

type userQuery struct {
	db *sql.DB
}

// Save создание пользователя
func (q *userQuery) Save(user dto.UserCreate) (id int, err error) {
	query := "INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"

	err = q.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, errors.New("пользователь уже существует")
		}
		return 0, err
	}

	return id, nil
}

// GetUserByEmail получить пользователя по email
func (q *userQuery) GetUserByEmail(email string) (user model.User, err error) {
	query := "SELECT id, password_hash FROM users WHERE email = $1"

	err = q.db.QueryRow(query, email).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("не верный логин или пароль")
		}
		return user, err
	}

	return user, nil
}
