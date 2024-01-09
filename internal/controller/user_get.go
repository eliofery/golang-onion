package controller

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// GetUser получение данных пользователя
func (c *ServiceController) GetUser(ctx fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || userId == 0 {
		userId = *c.authService.GetUserIdFromToken(ctx)
	}

	user, err := c.userService.GetById(userId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "данные пользователя",
		"user":    user,
	})
}
