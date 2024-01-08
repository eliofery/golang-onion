package route

import (
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func (r *Router) userRoute(app *fiber.App) {
	api := app.Group(apiV1)

	user := api.Group("/user", middleware.IsAuth)
	user.Post("/create", r.handler.CreateUser)
}
