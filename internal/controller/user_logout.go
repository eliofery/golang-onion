package controller

import (
	"github.com/gofiber/fiber/v3"
)

// Logout выход пользователя из системы
func (c *ServiceController) Logout(ctx fiber.Ctx) error {
	userId := c.authService.GetUserIdFromToken(ctx)

	if err := c.authService.Logout(ctx, userId); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "пользователь вышел из системы",
	})
}
