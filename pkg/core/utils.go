package core

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

var (
	ErrServerConnect = errors.New("ошибка при запуске сервера")
)

const (
	urlDefault  = "127.0.0.1"
	portDefault = "3000"
)

// RegisterMiddlewares регистрация промежуточного программного обеспечения
func (a *App) registerMiddlewares(app *fiber.App, middlewares []fiber.Handler) {
	for _, middleware := range middlewares {
		app.Use(middleware)
	}
}

// RegisterRoutes регистрация маршрутов
func (a *App) registerRoutes(app *fiber.App, routes []Route) {
	for _, route := range routes {
		route.Setup(app)
	}
}

// MustListen регистрация маршрутов
func (a *App) listen(app *fiber.App) error {
	op := "app.listen"

	port := a.config.Get("SERVER_PORT")
	if port == "" {
		port = portDefault
	}

	url := a.config.Get("SERVER_URL")
	if url == "" {
		url = urlDefault
	}

	if err := app.Listen(fmt.Sprintf("%s:%s", url, port)); err != nil {
		log.Error(fmt.Errorf("%s: %w", op, ErrServerConnect))
		return err
	}

	return nil
}
