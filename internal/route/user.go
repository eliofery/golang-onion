package route

import (
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func (r *Router) userRoute(app *fiber.App) {
	api := app.Group(apiV1)

	api.Get("/users", r.handler.GetUserAll, middleware.IsAuth)

	user := api.Group("/user", middleware.IsAuth)
	user.Get("/", r.handler.GetUser)
	user.Get("/:id", r.handler.GetUser)
	user.Put("/update", r.handler.UpdateUser)
	user.Put("/:id/update", r.handler.UpdateUser)
	user.Post("/create", r.handler.CreateUser)
}
