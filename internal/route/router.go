package route

import (
	"github.com/eliofery/golang-onion/internal/controller"
	"github.com/eliofery/golang-onion/pkg/core"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

const (
	apiV1 = "/api/v1"
)

// Router маршрутизатор
type Router struct {
	handler controller.ServiceController
}

func NewRouter(handler controller.ServiceController) core.Route {
	log.Info("регистрация маршрутов")
	return &Router{handler: handler}
}

// Register регистрация маршрутов
func (r *Router) Register(app *fiber.App) {
	r.userRoute(app)
	r.authRoute(app)
	r.roleRoute(app)
	r.permissionRoute(app)
}
