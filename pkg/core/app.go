package core

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"time"
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
func (a *App) MustRun(ctx context.Context) {
	log.Info("запуск приложения")
	op := "app.MustRun"

	defer func() {
		if err := a.db.Close(); err != nil {
			log.Errorf("%s: %s", op, err)
		}
	}()

	if err := database.Migrate(a.db); err != nil {
		log.Errorf("%s: %s", op, err)
	}

	server := fiber.New(a.options)
	a.registerMiddlewares(server, a.middlewares)
	a.registerRoutes(server, a.routes)

	ch := make(chan error, 1)
	go func() {
		if err := a.listen(server); err != nil {
			ch <- fmt.Errorf("не удалось запустить сервер: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		log.Fatalf("%s: %s", op, err)
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.ShutdownWithContext(timeout); err != nil {
			log.Errorf("%s: %s", op, err)
		}
	}

	log.Info("остановка приложения")
}
