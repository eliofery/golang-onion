package route

import (
	"github.com/eliofery/golang-onion/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

// authRoute маршруты связанные с авторизацией
func (r *Router) authRoute(app *fiber.App) {
	api := app.Group(apiV1)

	api.Post("/signup", r.handler.SignUp, middleware.IsGuest)
	api.Post("/signin", r.handler.SignIn, middleware.IsGuest)
	api.Post("/logout", r.handler.Logout, middleware.IsAuth)
}
