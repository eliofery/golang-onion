package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v3/log"
)

// DAO предоставляет доступ к механизму взаимодействия с данными в базе данных
type DAO interface {
	NewUserQuery() UserQuery
}

type dao struct {
	db *sql.DB
}

func NewDAO(db *sql.DB) DAO {
	log.Info("инициализация DAO")
	return &dao{db: db}
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{db: d.db}
}
