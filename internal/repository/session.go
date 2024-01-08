package repository

import (
	"database/sql"
)

// SessionQuery содержит запросы в базу данных для манипуляции с пользователями
type SessionQuery interface {
	Save(userId int, token string) error
	DeleteByToken(token string) error
	DeleteByUserId(userId int) error
	VerifyToken(token string) error
}

type sessionQuery struct {
	db *sql.DB
}

// Save сохранение токена
func (q *sessionQuery) Save(userId int, token string) error {
	query := "INSERT INTO sessions (token, user_id) VALUES ($1, $2)"
	_, err := q.db.Exec(query, token, userId)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByToken удаление сессии по токену
func (q *sessionQuery) DeleteByToken(token string) error {
	query := "DELETE FROM sessions WHERE token = $1"
	_, err := q.db.Exec(query, token)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByUserId удаление сессии по id пользователя
func (q *sessionQuery) DeleteByUserId(userId int) error {
	query := "DELETE FROM sessions WHERE user_id = $1"
	_, err := q.db.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

// VerifyToken проверка токена на наличие в БД
func (q *sessionQuery) VerifyToken(token string) error {
	query := "SELECT id FROM sessions WHERE token = $1"
	_, err := q.db.Exec(query, token)
	if err != nil {
		return err
	}

	return nil
}
