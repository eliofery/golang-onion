package controller

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
)

// CreateRole создание роли
func (c *ServiceController) CreateRole(ctx fiber.Ctx) error {
	var role dto.RolePermission
	if err := c.bodyValidate(ctx, &role); err != nil {
		return err
	}

	createdRole, err := c.roleService.Create(role)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "роль создана",
		"role":    createdRole,
	})
}
