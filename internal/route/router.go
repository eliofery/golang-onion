package route

import (
	"github.com/eliofery/golang-angular/internal/controller"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/gofiber/fiber/v3"
)

// Router маршрутизатор
type Router struct {
	handler controller.ServiceController
}

func NewRouter(handler controller.ServiceController) core.Route {
	return &Router{handler: handler}
}

func (r *Router) Setup(app *fiber.App) {
	r.userRoute(app)
	r.authRoute(app)
}
