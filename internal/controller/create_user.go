package controller

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// CreateUser создание пользователя
func (c *ServiceController) CreateUser(ctx fiber.Ctx) error {
	// TODO: убрать
	userId := c.authService.GetUserIdFromToken(ctx)
	log.Info(userId)

	var user dto.UserCreate
	if err := c.bodyValidate(ctx, &user); err != nil {
		return err
	}

	id, err := c.authService.Register(user)
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
