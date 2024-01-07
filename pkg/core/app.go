package core

import (
	"database/sql"
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
	log.Info("инициализация приложения")
	return &App{
		config: conf,
		db:     db,
	}
}

// SetOptions задает дополнительные настройки для сервера
func (a *App) SetOptions(options fiber.Config) *App {
	log.Info("настройка fiber.Config")
	a.options = options

	return a
}

// UseMiddlewares использование промежуточное программное обеспечение
func (a *App) UseMiddlewares(injections ...any) *App {
	log.Info("инициализация middlewares")

	for _, injection := range injections {
		if middleware, ok := injection.(fiber.Handler); ok {
			a.middlewares = append(a.middlewares, middleware)
		}
	}

	return a
}

// UseRoutes использование маршрутов
func (a *App) UseRoutes(injections ...Route) *App {
	log.Info("инициализация маршрутов")

	for _, route := range injections {
		a.routes = append(a.routes, route)
	}

	return a
}

// MustRun запуск приложения с обработкой ошибок
func (a *App) MustRun() {
	log.Info("запуск приложения")
	op := "app.MustRun"

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Fatalf("%s: %s", op, err)
		}
	}(a.db)

	if err := database.Migrate(a.db); err != nil {
		log.Fatalf("%s: %s", op, err)
	}

	server := fiber.New(a.options)
	a.registerMiddlewares(server, a.middlewares)
	a.registerRoutes(server, a.routes)
	if err := a.listen(server); err != nil {
		log.Fatalf("%s: %s", op, err)
	}
}
