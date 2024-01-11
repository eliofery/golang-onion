package controller

import (
	"github.com/gofiber/fiber/v3"
)

// DeleteRole удаление данных роли
func (c *ServiceController) DeleteRole(ctx fiber.Ctx) error {
	roleId, err := c.getIdValidate(ctx)
	if err != nil {
		return err
	}

	if err = c.roleService.Delete(*roleId); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "роль удалена",
		"role":    roleId,
	})
}
