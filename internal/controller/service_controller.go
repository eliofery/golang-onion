package controller

import (
	"github.com/eliofery/golang-angular/internal/service"
	"github.com/eliofery/golang-angular/pkg/utils"
)

// ServiceController обработчик маршрутов
type ServiceController struct {
	jwtService  utils.Jwt
	authService service.AuthService
	userService service.UserService
}

func NewController(
	jwtService utils.Jwt,
	authService service.AuthService,
	userService service.UserService,
) ServiceController {
	return ServiceController{
		jwtService:  jwtService,
		authService: authService,
		userService: userService,
	}
}
