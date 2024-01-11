package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/eliofery/golang-onion/internal/model"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

// RoleQuery содержит запросы в базу данных для манипуляции с ролями
type RoleQuery interface {
	GetAll(limit, offset int) (roles []model.Role, err error)
	GetTotalCount() (count int, err error)
	GetById(roleId int) (role *model.Role, err error)
	Update(role model.Role) (updateRole *model.Role, err error)
	Delete(roleId int) error
	Create(user model.Role) (role *model.Role, err error)
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
func (q *roleQuery) Update(role model.Role) (*model.Role, error) {
	query := "UPDATE roles SET name = $1 WHERE id = $2"
	_, err := q.db.Exec(query, role.Name, role.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	query = "DELETE FROM role_permissions WHERE role_id = $1"
	_, err = q.db.Exec(query, role.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var args []any
	args = append(args, role.ID)

	placeholders := make([]string, len(role.Permissions))
	for i, permission := range role.Permissions {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args = append(args, permission.ID)
	}

	permissionsString := fmt.Sprintf("ARRAY[%s]::int[]", strings.Join(placeholders, ","))

	query = "INSERT INTO role_permissions (role_id, permission_id) SELECT $1, unnest(" + permissionsString + ")"
	log.Info(query, args)
	_, err = q.db.Exec(query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &role, nil
}

// Delete удаление данных роли
func (q *roleQuery) Delete(roleId int) error {
	query := "DELETE FROM roles WHERE id = $1 RETURNING id"
	result, err := q.db.Exec(query, roleId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("роль не найдена")
	}

	return nil
}

// Create создание пользователя
func (q *roleQuery) Create(role model.Role) (*model.Role, error) {
	args := make([]any, 0, len(role.Permissions)+1)
	args = append(args, role.Name)
	for _, permission := range role.Permissions {
		args = append(args, permission.ID)
	}

	placeholders := make([]string, len(role.Permissions))
	for i := range role.Permissions {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
	}
	permissionsString := fmt.Sprintf("ARRAY[%s]::int[]", strings.Join(placeholders, ","))

	query := `WITH inserted_role AS (
        INSERT INTO roles(name) VALUES($1) RETURNING id
    )
    INSERT INTO role_permissions(role_id, permission_id)
    SELECT (SELECT id FROM inserted_role), unnest(` + permissionsString + `) AS permission_id RETURNING role_id`

	if err := q.db.QueryRow(query, args...).Scan(&role.ID); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			return nil, errors.New("роль уже существует")
		}

		return nil, err
	}

	return &role, nil
}
