package main

import (
	"context"
	"github.com/eliofery/golang-angular/internal/controller"
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/internal/route"
	"github.com/eliofery/golang-angular/internal/service"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/eliofery/golang-angular/pkg/database/postgres"
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"os"
	"os/signal"
)

func main() {
	// Модули приложения
	conf := config.MustInit(godotenv.New())
	db := database.MustConnect(postgres.New(conf))
	validate := utils.NewValidate(validator.New())

	// Логика приложения
	dao := repository.NewDAO(db)
	handler := controller.NewServiceController(
		model.NewValidate(validate),
		utils.NewJwt(conf),
		service.NewAuthService(dao),
		service.NewUserService(dao),
	)

	// Запуск приложения
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	core.New(conf, db).
		SetOptions(fiber.Config{
			ErrorHandler: middleware.ErrorHandler,
		}).
		UseMiddlewares(
			middleware.Cors(conf),
		).
		UseRoutes(
			route.NewRouter(handler),
		).
		MustRun(ctx)
}
