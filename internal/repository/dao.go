package repository

import "database/sql"

// DAO предоставляет доступ к механизму взаимодействия с данными в базе данных
type DAO interface{}

type dao struct {
	db *sql.DB
}

func NewDAO(db *sql.DB) DAO {
	return &dao{db: db}
}
