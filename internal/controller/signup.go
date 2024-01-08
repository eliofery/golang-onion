package controller

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
)

// SignUp вход в систему нового пользователя
func (c *ServiceController) SignUp(ctx fiber.Ctx) error {
	var user dto.UserCreate
	if err := c.bodyValidate(ctx, &user); err != nil {
		return err
	}

	token, err := c.authService.RegisterAndAuth(ctx, user)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "пользователь создан и авторизован",
		"token":   token,
	})
}
