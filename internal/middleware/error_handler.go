package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// ErrorHandler обработчик ошибок
func ErrorHandler(ctx fiber.Ctx, err error) error {
	var (
		message string

		pgErr   *pgconn.PgError
		jsonErr *json.SyntaxError
	)

	log.Errorf("%T: %s", err, err)

	switch {
	case errors.As(err, &jsonErr):
		message = "некорректный json формат"
	case errors.As(err, &pgErr):
		if pgErr.Code == pgerrcode.UniqueViolation {
			message = "пользователь уже существует"
		} else {
			message = "ошибка базы данных"
		}
	default:
		message = err.Error()
	}

	return ctx.JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
