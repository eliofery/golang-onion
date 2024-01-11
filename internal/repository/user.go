package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"strings"
)

// UserQuery содержит запросы в базу данных для манипуляции с пользователями
type UserQuery interface {
	Create(user dto.UserCreate) (userId int, err error)
	GetByEmail(email string) (user *model.User, err error)
	GetById(userId int) (user *model.User, err error)
	GetAll(limit, offset int) ([]model.User, error)
	GetTotalCount() (int, error)
	Update(user dto.UserUpdate) (*model.User, error)
	Delete(userId int) error
}

type userQuery struct {
	db *sql.DB
}

// Create создание пользователя
func (q *userQuery) Create(user dto.UserCreate) (int, error) {
	var userId int

	query := "INSERT INTO users (first_name, last_name, email, password_hash, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := q.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.RoleID).Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return 0, errors.New("пользователь уже существует")
			case pgerrcode.ForeignKeyViolation:
				return 0, errors.New("роль не найдена")
			}
		}
		return 0, err
	}

	return userId, nil
}

// GetByEmail получить пользователя по email
func (q *userQuery) GetByEmail(email string) (*model.User, error) {
	var user model.User

	query := "SELECT id, password_hash FROM users WHERE email = $1"
	err := q.db.QueryRow(query, email).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("не верный логин или пароль")
		}
		return nil, err
	}

	return &user, nil
}

// GetById получить пользователя по id
func (q *userQuery) GetById(userId int) (*model.User, error) {
	var user model.User

	query := `SELECT users.id, first_name, last_name, email, roles.id, roles.name
    FROM users
    INNER JOIN roles ON users.role_id = roles.id
    WHERE users.id = $1`
	err := q.db.QueryRow(query, userId).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role.ID, &user.Role.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}

	return &user, nil
}

// GetAll получить всех пользователей
func (q *userQuery) GetAll(limit, offset int) ([]model.User, error) {
	query := `SELECT users.id, first_name, last_name, email, roles.id, roles.name
    FROM users
    INNER JOIN roles ON users.role_id = roles.id
    LIMIT $1
    OFFSET $2`
	rows, err := q.db.Query(query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("пользователи не найдены")
		}
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role.ID, &user.Role.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetTotalCount получить общее количество пользователей
func (q *userQuery) GetTotalCount() (int, error) {
	query := "SELECT COUNT(*) FROM users"

	var count int
	err := q.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Update обновление данных пользователя
func (q *userQuery) Update(user dto.UserUpdate) (*model.User, error) {
	var (
		setClauses []string
		args       []any
	)

	if user.FirstName != "" {
		setClauses = append(setClauses, fmt.Sprintf("first_name = $%d", len(args)+1))
		args = append(args, user.FirstName)
	}

	if user.LastName != "" {
		setClauses = append(setClauses, fmt.Sprintf("last_name = $%d", len(args)+1))
		args = append(args, user.LastName)
	}

	if user.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, user.Email)
	}

	if user.Password != "" {
		setClauses = append(setClauses, fmt.Sprintf("password_hash = $%d", len(args)+1))
		args = append(args, user.Password)
	}

	if user.RoleID != 0 {
		setClauses = append(setClauses, fmt.Sprintf("role_id = $%d", len(args)+1))
		args = append(args, user.RoleID)
	}

	if len(setClauses) == 0 {
		return nil, errors.New("нет данных для обновления")
	}

	setClause := strings.Join(setClauses, ", ")
	query := fmt.Sprintf("UPDATE users "+
		"SET %s "+
		"FROM roles "+
		"WHERE users.id = $%d AND users.role_id = roles.id "+
		"RETURNING users.id, first_name, last_name, email, roles.id, roles.name",
		setClause, len(args)+1)
	args = append(args, user.ID)

	var updateUser model.User
	err := q.db.QueryRow(query, args...).Scan(&updateUser.ID, &updateUser.FirstName, &updateUser.LastName, &updateUser.Email, &updateUser.Role.ID, &updateUser.Role.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, errors.New("пользователь с таким email уже существует")
		}
		return nil, err
	}

	return &updateUser, nil
}

// Delete удаление данных пользователя
func (q *userQuery) Delete(userId int) error {
	query := "DELETE FROM users WHERE id = $1 RETURNING id"
	result, err := q.db.Exec(query, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("пользователь не найден")
	}

	return nil
}
