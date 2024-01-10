package controller

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// GetRoles получение ролей
func (c *ServiceController) GetRoles(ctx fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	data, err := c.roleService.GetAll(page)
	if err != nil {
		return err
	}

	resp := fiber.Map{
		"success": true,
		"message": "список ролей",
		"users":   data.Roles,
		"meta":    data.Meta,
	}

	if data.Roles == nil {
		resp = fiber.Map{
			"success": false,
			"message": "роли не найдены",
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
