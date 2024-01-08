package repository

import (
	"database/sql"
)

// SessionQuery содержит запросы в базу данных для манипуляции с пользователями
type SessionQuery interface {
	Save(token string) error
	Delete(token string) error
}

type sessionQuery struct {
	db *sql.DB
}

// Save сохранение токена
func (q *sessionQuery) Save(token string) error {
	query := "INSERT INTO sessions (token) VALUES ($1)"
	_, err := q.db.Exec(query, token)
	if err != nil {
		return err
	}

	return nil
}

// Delete удаление токена
func (q *sessionQuery) Delete(token string) error {
	query := "DELETE FROM sessions WHERE token = $1"
	_, err := q.db.Exec(query, token)
	if err != nil {
		return err
	}

	return nil
}
