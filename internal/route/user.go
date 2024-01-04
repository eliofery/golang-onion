package route

import (
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/gofiber/fiber/v3"
)

// User маршруты связанные с авторизацией пользователя
type User struct {
}

func NewUser() core.Route {
	return &User{}
}

func (a *User) Setup(app *fiber.App) {
	auth := app.Group("/api/v1/user")
	auth.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Пользователь",
		})
	})
}
