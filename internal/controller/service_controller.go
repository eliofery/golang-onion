package controller

import (
	"errors"
	"github.com/eliofery/golang-angular/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// ServiceController обработчик маршрутов
type ServiceController struct {
	validateService service.ValidateService
	authService     service.AuthService
	userService     service.UserService
	roleService     service.RoleService
}

func NewServiceController(
	validateService service.ValidateService,
	authService service.AuthService,
	userService service.UserService,
	roleService service.RoleService,
) ServiceController {
	log.Info("инициализация сервисов контроллера")

	return ServiceController{
		validateService: validateService,
		authService:     authService,
		userService:     userService,
		roleService:     roleService,
	}
}

// bodyValidate валидация входных данных
func (c *ServiceController) bodyValidate(ctx fiber.Ctx, data any) error {
	if err := ctx.Bind().Body(&data); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return errors.New("некорректный json формат")
	}

	if errMessages := c.validateService.ValidateData(data); len(errMessages) > 0 {
		ctx.Status(fiber.StatusBadRequest)
		return errors.Join(errMessages...)
	}

	return nil
}
