package controller

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// UserDelete удаление данных пользователя
func (c *ServiceController) UserDelete(ctx fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || userId <= 0 {
		userId = *c.authService.GetUserIdFromToken(ctx)
	}

	if err = c.userService.Delete(ctx, userId); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "пользователь удален",
		"user":    userId,
	})
}
