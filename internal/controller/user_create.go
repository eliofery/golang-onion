package controller

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
)

// CreateUser создание пользователя
func (c *ServiceController) CreateUser(ctx fiber.Ctx) error {
	var user dto.UserCreate
	if err := c.bodyValidate(ctx, &user); err != nil {
		return err
	}

	id, err := c.userService.Create(user)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "пользователь создан",
		"id":      id,
	})
}
