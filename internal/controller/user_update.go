package controller

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// UpdateUser обновление данных пользователя
func (c *ServiceController) UpdateUser(ctx fiber.Ctx) error {
	var user dto.UserUpdate
	if err := c.bodyValidate(ctx, &user); err != nil {
		return err
	}

	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || userId <= 0 {
		userId = *c.authService.GetUserIdFromToken(ctx)
	}

	user.ID = userId
	updateUser, err := c.userService.Update(user)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "пользователь обновлен",
		"user":    updateUser,
	})
}
