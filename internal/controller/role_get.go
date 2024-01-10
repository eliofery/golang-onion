package controller

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// GetRole получение данных роли
func (c *ServiceController) GetRole(ctx fiber.Ctx) error {
	roleId, err := strconv.Atoi(ctx.Params("id", "1"))
	if err != nil || roleId <= 0 {
		ctx.Status(fiber.StatusBadRequest)
		return errors.New("некорректный идентификатор роли")
	}

	role, err := c.roleService.GetById(roleId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "данные роли",
		"role":    role,
	})
}
