package route

import "github.com/gofiber/fiber/v3"

// authRoute маршруты связанные с авторизацией
func (r *Router) authRoute(app *fiber.App) {
	api := app.Group(apiV1)

	api.Post("/signup", r.handler.SignUp)
	api.Post("/signin", r.handler.SignIn)
}
