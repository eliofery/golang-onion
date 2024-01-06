package main

import (
	"github.com/eliofery/golang-angular/internal/controller"
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/internal/route"
	"github.com/eliofery/golang-angular/internal/service"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/eliofery/golang-angular/pkg/database/postgres"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Первостепенная инициализация
	conf := config.MustInit(godotenv.New())
	db := database.MustConnect(postgres.New(conf))

	// Связывание логики приложения
	dao := repository.NewDAO(db)
	handler := controller.NewController(
		service.NewAuthService(dao),
		service.NewUserService(dao),
	)

	// Запуск приложения
	core.New(conf, db).
		SetOptions(fiber.Config{
			ErrorHandler: middleware.NotFound,
		}).
		UseMiddlewares(
			middleware.Cors(conf),
		).
		UseRoutes(
			route.NewRouter(handler),
		).
		MustRun()
}
