package repository

import (
	"database/sql"
	"errors"
	"github.com/eliofery/golang-angular/internal/model"
)

// RoleQuery содержит запросы в базу данных для манипуляции с ролями
type RoleQuery interface {
	GetAll(limit, offset int) (roles []model.Role, err error)
	GetTotalCount() (count int, err error)
}

type roleQuery struct {
	db *sql.DB
}

// GetAll получить всех ролей
func (q *roleQuery) GetAll(limit, offset int) ([]model.Role, error) {
	query := "SELECT id, name FROM roles LIMIT $1 OFFSET $2"
	rows, err := q.db.Query(query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("роли не найдены")
		}
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role
	for rows.Next() {
		var role model.Role
		if err = rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

// GetTotalCount получить общее количество ролей
func (q *roleQuery) GetTotalCount() (int, error) {
	query := "SELECT COUNT(*) FROM roles"

	var count int
	err := q.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
