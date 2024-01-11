package route

import (
	"github.com/eliofery/golang-onion/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func (r *Router) permissionRoute(app *fiber.App) {
	api := app.Group(apiV1)

	api.Get("/permissions", r.handler.GetPermissions, middleware.IsAuth)
}
