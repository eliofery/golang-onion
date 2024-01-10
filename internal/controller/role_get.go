package controller

import (
	"github.com/gofiber/fiber/v3"
)

// GetRole получение данных роли
func (c *ServiceController) GetRole(ctx fiber.Ctx) error {
	roleId, err := c.idValidate(ctx)
	if err != nil {
		return err
	}

	role, err := c.roleService.GetById(*roleId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "данные роли",
		"role":    role,
	})
}
