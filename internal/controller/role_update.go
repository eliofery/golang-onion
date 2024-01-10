package controller

import (
	"errors"
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// UpdateRole обновление данных роли
func (c *ServiceController) UpdateRole(ctx fiber.Ctx) error {
	var role dto.Role
	if err := c.bodyValidate(ctx, &role); err != nil {
		return err
	}

	roleId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || roleId <= 0 {
		ctx.Status(fiber.StatusBadRequest)
		return errors.New("некорректный идентификатор роли")
	}

	role.ID = roleId
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
