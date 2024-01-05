package route

import (
	"github.com/eliofery/golang-angular/internal/controller"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/gofiber/fiber/v3"
)

// Auth маршруты связанные с авторизацией пользователя
type Auth struct {
	handler controller.AuthController
}

func NewAuth(handler controller.AuthController) core.Route {
	return &Auth{handler: handler}
}

func (a *Auth) Setup(app *fiber.App) {
	api := app.Group("/api/v1/auth")
	api.Get("/", a.handler.Auth)
}
