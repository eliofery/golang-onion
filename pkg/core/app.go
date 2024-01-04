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
	routes      []Route
	middlewares []fiber.Handler
}

func New(conf config.Config, db *sql.DB) *App {
	return &App{
		config: conf,
		db:     db,
	}
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

	server := fiber.New()
	a.registerMiddlewares(server, a.middlewares)
	a.registerRoutes(server, a.routes)
	if err := a.listen(server); err != nil {
		log.Fatal(fmt.Errorf("%s: %w", op, err))
	}
}
