package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v3/log"
)

// DAO предоставляет доступ к механизму взаимодействия с данными в базе данных
type DAO interface {
	NewUserQuery() UserQuery
	NewSessionQuery() SessionQuery
	NewRoleQuery() RoleQuery
	NewPermissionQuery() PermissionQuery
}

type dao struct {
	db *sql.DB
}

func NewDAO(db *sql.DB) DAO {
	log.Info("инициализация DAO")
	return &dao{db: db}
}

// NewUserQuery запросы в базу данных для пользователей
func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{db: d.db}
}

// NewSessionQuery запросы в базу данных для сессий
func (d *dao) NewSessionQuery() SessionQuery {
	return &sessionQuery{db: d.db}
}

// NewRoleQuery запросы в базу данных для ролей
func (d *dao) NewRoleQuery() RoleQuery {
	return &roleQuery{db: d.db}
}

// NewPermissionQuery запросы в базу данных для разрешений
func (d *dao) NewPermissionQuery() PermissionQuery {
	return &permissionQuery{db: d.db}
}
