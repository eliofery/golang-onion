package main

import (
	"context"
	"github.com/eliofery/golang-onion/internal/controller"
	"github.com/eliofery/golang-onion/internal/middleware"
	"github.com/eliofery/golang-onion/internal/repository"
	"github.com/eliofery/golang-onion/internal/route"
	"github.com/eliofery/golang-onion/internal/service"
	"github.com/eliofery/golang-onion/internal/validation"
	"github.com/eliofery/golang-onion/pkg/config"
	"github.com/eliofery/golang-onion/pkg/config/godotenv"
	"github.com/eliofery/golang-onion/pkg/core"
	"github.com/eliofery/golang-onion/pkg/database"
	"github.com/eliofery/golang-onion/pkg/database/postgres"
	"github.com/eliofery/golang-onion/pkg/utils"
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
	jwt := utils.NewTokenManager(conf)

	// Логика приложения
	dao := repository.NewDAO(db)
	handler := controller.NewServiceController(
		service.NewValidateService(validate).Register(
			validation.TestValidate(),
		),
		service.NewAuthService(dao, jwt),
		service.NewUserService(dao, conf, jwt),
		service.NewRoleService(dao, conf),
		service.NewPermissionService(dao, conf),
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
			middleware.SetUserIdFromToken(dao, jwt),
		).
		UseRoutes(
			route.NewRouter(handler),
		).
		MustRun(ctx)
}
