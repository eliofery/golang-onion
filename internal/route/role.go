package route

import (
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func (r *Router) roleRoute(app *fiber.App) {
	api := app.Group(apiV1)

	api.Get("/roles", r.handler.GetRoles, middleware.IsAuth)

	user := api.Group("/role", middleware.IsAuth)
	user.Get("/:id", r.handler.GetRole)
	//user.Put("/update", r.handler.UpdateUser)
	//user.Put("/:id/update", r.handler.UpdateUser)
	//user.Delete("/delete", r.handler.UserDelete)
	//user.Delete("/:id/delete", r.handler.UserDelete)
	//user.Post("/create", r.handler.CreateUser)
}
