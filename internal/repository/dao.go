package repository

import "database/sql"

// DAO предоставляет доступ к механизму взаимодействия с данными в базе данных
type DAO interface {
	NewUserQuery() UserQuery
}

type dao struct {
	db *sql.DB
}

func NewDAO(db *sql.DB) DAO {
	return &dao{db: db}
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{db: d.db}
}
