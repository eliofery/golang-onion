package controller

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// GetPermissions получение всех разрешений
func (c *ServiceController) GetPermissions(ctx fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	data, err := c.permissionService.GetAll(page)
	if err != nil {
		return err
	}

	resp := fiber.Map{
		"success":     true,
		"message":     "список разрешений",
		"permissions": data.Permissions,
		"meta":        data.Meta,
	}

	if data.Permissions == nil {
		resp = fiber.Map{
			"success": false,
			"message": "разрешения не найдены",
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
