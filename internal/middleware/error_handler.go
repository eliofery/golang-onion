package middleware

import (
	"github.com/gofiber/fiber/v3"
)

// ErrorHandler обработчик ошибок
func ErrorHandler(ctx fiber.Ctx, err error) error {
	return ctx.JSON(fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}
