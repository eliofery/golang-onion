package controller

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// DeleteUser удаление данных пользователя
func (c *ServiceController) DeleteUser(ctx fiber.Ctx) error {
	// TODO: вынести удаление авторизованного пользователя отдельно
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
