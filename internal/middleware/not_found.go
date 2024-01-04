package middleware

import "github.com/gofiber/fiber/v3"

// NotFound обработчик ошибки 404
func NotFound(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "ссылка не найдена",
	})
}
