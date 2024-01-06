package controller

import (
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/service"
	"github.com/eliofery/golang-angular/pkg/utils"
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
