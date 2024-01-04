package route

import (
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/gofiber/fiber/v3"
)

// Auth маршруты связанные с авторизацией пользователя
type Auth struct{}

func NewAuth() core.Route {
	return &Auth{}
}

func (a *Auth) Setup(app *fiber.App) {
	api := app.Group("/api/v1/auth")
	api.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Авторизация",
		})
	})
}
