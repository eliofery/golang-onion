package repository

import (
	"database/sql"
	"errors"
	"github.com/eliofery/golang-angular/internal/model"
)

// PermissionQuery содержит запросы в базу данных для манипуляции с разрешениями
type PermissionQuery interface {
	GetAll(limit, offset int) (permissions []model.Permission, err error)
	GetTotalCount() (int, error)
}

type permissionQuery struct {
	db *sql.DB
}

// GetAll получить всех разрешений
func (q *permissionQuery) GetAll(limit, offset int) ([]model.Permission, error) {
	query := "SELECT id, name, description FROM permissions LIMIT $1 OFFSET $2"
	rows, err := q.db.Query(query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("разрешения не найдены")
		}
		return nil, err
	}
	defer rows.Close()

	var permissions []model.Permission
	for rows.Next() {
		var permission model.Permission
		if err = rows.Scan(&permission.ID, &permission.Name, &permission.Description); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

// GetTotalCount получить общее количество разрешений
func (q *permissionQuery) GetTotalCount() (int, error) {
	query := "SELECT COUNT(*) FROM permissions"

	var count int
	err := q.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
