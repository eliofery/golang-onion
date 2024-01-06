package controller

import (
	"github.com/gofiber/fiber/v3"
)

// CreateUser создание пользователя
func (s *ServiceController) CreateUser(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Регистрация",
	})
}
