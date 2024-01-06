package route

import "github.com/gofiber/fiber/v3"

func (r *Router) userRoute(app *fiber.App) {
	auth := app.Group("/api/v1/user")
	auth.Get("/", r.handler.GetUser)
}
