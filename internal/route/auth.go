package route

import "github.com/gofiber/fiber/v3"

func (r *Router) authRoute(app *fiber.App) {
	api := app.Group("/api/v1/auth")
	api.Get("/", r.handler.Auth)
}
