package core

import (
	"database/sql"
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// App приложение
type App struct {
	config      config.Config
	db          *sql.DB
	options     fiber.Config
	middlewares []fiber.Handler
	routes      []Route
}

func New(conf config.Config, db *sql.DB) *App {
	return &App{
		config: conf,
		db:     db,
	}
}

// SetOptions задает дополнительные настройки для сервера
func (a *App) SetOptions(options fiber.Config) *App {
	a.options = options

	return a
}

// UseMiddlewares использование промежуточное программное обеспечение
func (a *App) UseMiddlewares(injections ...any) *App {
	for _, injection := range injections {
		if middleware, ok := injection.(fiber.Handler); ok {
			a.middlewares = append(a.middlewares, middleware)
		}
	}

	return a
}

// UseRoutes использование маршрутов
func (a *App) UseRoutes(injections ...Route) *App {
	for _, route := range injections {
		a.routes = append(a.routes, route)
	}

	return a
}

// MustRun запуск приложения с обработкой ошибок
func (a *App) MustRun() {
	op := "app.MustRun"

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Fatal(fmt.Errorf("%s: %w", op, err))
		}
	}(a.db)

	if err := database.Migrate(a.db); err != nil {
		log.Fatal(fmt.Errorf("%s: %w", op, err))
	}

	server := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "ссылка не найдена",
			})
		},
	})
	a.registerMiddlewares(server, a.middlewares)
	a.registerRoutes(server, a.routes)
	if err := a.listen(server); err != nil {
		log.Fatal(fmt.Errorf("%s: %w", op, err))
	}
}
