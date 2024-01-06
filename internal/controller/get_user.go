package controller

import "github.com/gofiber/fiber/v3"

// GetUser получение данных пользователя
func (s *ServiceController) GetUser(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Пользователь",
	})
}
