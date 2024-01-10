package controller

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// GetUsers получение данных пользователя
func (c *ServiceController) GetUsers(ctx fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	data, err := c.userService.GetAll(page)
	if err != nil {
		return err
	}

	resp := fiber.Map{
		"success": true,
		"message": "список пользователей",
		"users":   data.Users,
		"meta":    data.Meta,
	}

	if data.Users == nil {
		resp = fiber.Map{
			"success": false,
			"message": "пользователи не найдены",
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
