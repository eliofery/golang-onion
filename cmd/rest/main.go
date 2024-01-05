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
	authRoutes := route.NewAuth(controller.AuthController{
		UserService: service.NewUserService(dao),
	})
	userRoutes := route.NewUser(controller.UserController{
		AuthService: service.NewAuthService(dao),
	})

	// Запуск приложения
	rest := core.New(conf, db)
	rest.SetOptions(fiber.Config{
		ErrorHandler: middleware.NotFound,
	})
	rest.UseMiddlewares(
		middleware.Cors(conf),
	)
	rest.UseRoutes(
		authRoutes,
		userRoutes,
	)
	rest.MustRun()
}
