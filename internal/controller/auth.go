package controller

import "github.com/gofiber/fiber/v3"

// Auth авторизация пользователя
func (h *AuthController) Auth(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Авторизация",
	})
}
