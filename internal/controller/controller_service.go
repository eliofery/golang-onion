package controller

import "github.com/eliofery/golang-angular/internal/service"

// Обработчики контроллеров.
// Содержат методы обрабатывающие определенные запросы.

// AuthController авторизация пользователя
type AuthController struct {
	service.UserService
}

// UserController манипуляция с пользователями
type UserController struct {
	service.AuthService
}
