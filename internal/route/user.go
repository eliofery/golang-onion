package route

import "github.com/gofiber/fiber/v3"

func (r *Router) userRoute(app *fiber.App) {
	api := app.Group(apiV1)

	api.Post("/user/create", r.handler.CreateUser)
}
