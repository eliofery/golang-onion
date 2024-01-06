package route

import (
	"github.com/eliofery/golang-angular/internal/controller"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/gofiber/fiber/v3"
)

const (
	apiV1 = "/api/v1"
)

// Router маршрутизатор
type Router struct {
	handler controller.ServiceController
}

func NewRouter(handler controller.ServiceController) core.Route {
	return &Router{handler: handler}
}

// Register регистрация маршрутов
func (r *Router) Register(app *fiber.App) {
	r.userRoute(app)
	r.authRoute(app)
}
