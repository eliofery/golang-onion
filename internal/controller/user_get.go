package controller

import "github.com/gofiber/fiber/v3"

// GetUser получение данных пользователя
func (c *ServiceController) GetUser(ctx fiber.Ctx) error {
	userId := c.authService.GetUserIdFromToken(ctx)

	user, err := c.userService.GetUser(userId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Пользователь",
		"user":    user,
	})
}
