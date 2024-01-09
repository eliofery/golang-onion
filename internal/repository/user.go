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
	Save(user dto.UserCreate) (userId int, err error)
	GetUserByEmail(email string) (user *model.User, err error)
	GetUserById(userId int) (user *model.User, err error)
}

type userQuery struct {
	db *sql.DB
}

// Save создание пользователя
func (q *userQuery) Save(user dto.UserCreate) (int, error) {
	var userId int

	query := "INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"
	err := q.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password).Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, errors.New("пользователь уже существует")
		}
		return 0, err
	}

	return userId, nil
}

// GetUserByEmail получить пользователя по email
func (q *userQuery) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	query := "SELECT id, password_hash FROM users WHERE email = $1"
	err := q.db.QueryRow(query, email).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("не верный логин или пароль")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserById получить пользователя по id
func (q *userQuery) GetUserById(userId int) (*model.User, error) {
	var user model.User

	query := "SELECT id, first_name, last_name, email FROM users WHERE id = $1"
	err := q.db.QueryRow(query, userId).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}

	return &user, nil
}
