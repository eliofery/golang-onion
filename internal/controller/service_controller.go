package controller

import (
	"errors"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/service"
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

// ServiceController обработчик маршрутов
type ServiceController struct {
	validator model.Validate
	jwt       utils.Jwt

	authService service.AuthService
	userService service.UserService
}

func NewServiceController(
	validator model.Validate,
	jwt utils.Jwt,

	authService service.AuthService,
	userService service.UserService,
) ServiceController {
	return ServiceController{
		validator: validator,
		jwt:       jwt,

		authService: authService,
		userService: userService,
	}
}

// bodyValidate валидация входных данных
func (c *ServiceController) bodyValidate(ctx fiber.Ctx, data any) error {
	if err := ctx.Bind().Body(&data); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}

	if errMessages := c.validator.Validation(data); len(errMessages) > 0 {
		ctx.Status(fiber.StatusBadRequest)
		return errors.Join(errMessages...)
	}

	return nil
}
