package controller

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
)

// SignIn авторизация пользователя
func (c *ServiceController) SignIn(ctx fiber.Ctx) error {
	var user dto.UserAuth
	if err := c.bodyValidate(ctx, &user); err != nil {
		return err
	}

	token, err := c.authService.Auth(ctx, user)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "пользователь авторизован",
		"token":   token,
	})
}
