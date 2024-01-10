package repository

import (
	"database/sql"
	"errors"
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// RoleQuery содержит запросы в базу данных для манипуляции с ролями
type RoleQuery interface {
	GetAll(limit, offset int) (roles []model.Role, err error)
	GetTotalCount() (count int, err error)
	GetById(roleId int) (role *model.Role, err error)
	Update(role dto.Role) (updateRole *model.Role, err error)
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

// GetById получить роль по id
func (q *roleQuery) GetById(roleId int) (*model.Role, error) {
	var role model.Role

	query := "SELECT id, name FROM roles WHERE id = $1"
	err := q.db.QueryRow(query, roleId).Scan(&role.ID, &role.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("роль не найдена")
		}
		return nil, err
	}

	return &role, nil
}

// Update обновление данных ролей
func (q *roleQuery) Update(role dto.Role) (*model.Role, error) {
	query := "UPDATE roles SET name = $1 WHERE id = $2 RETURNING id, name"

	var updateRole model.Role
	err := q.db.QueryRow(query, role.Name, role.ID).Scan(&updateRole.ID, &updateRole.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, errors.New("роль с таким именем существует")
		} else if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("роль не найдена")
		}
		return nil, err
	}

	return &updateRole, nil
}
