package route

import (
	"github.com/eliofery/golang-angular/internal/controller"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/gofiber/fiber/v3"
)

// User маршруты связанные с пользователями
type User struct {
	handler controller.UserController
}

func NewUser(handler controller.UserController) core.Route {
	return &User{handler: handler}
}

func (u *User) Setup(app *fiber.App) {
	auth := app.Group("/api/v1/user")
	auth.Get("/", u.handler.GetUser)
}
