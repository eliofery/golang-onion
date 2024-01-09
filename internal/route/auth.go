package route

import (
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

// authRoute маршруты связанные с авторизацией
func (r *Router) authRoute(app *fiber.App) {
	api := app.Group(apiV1)

	api.Post("/logout", r.handler.Logout, middleware.IsAuth)

	guest := api.Group("/", middleware.IsGuest)
	guest.Post("/signup", r.handler.SignUp)
	guest.Post("/signin", r.handler.SignIn)
}
