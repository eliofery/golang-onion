package controller

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
)

// UpdateRole обновление данных роли
func (c *ServiceController) UpdateRole(ctx fiber.Ctx) error {
	var role dto.Role
	if err := c.bodyValidate(ctx, &role); err != nil {
		return err
	}

	roleId, err := c.idValidate(ctx)
	if err != nil {
		return err
	}

	role.ID = *roleId
	updateUser, err := c.roleService.Update(role)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "роль обновлен",
		"user":    updateUser,
	})
}
