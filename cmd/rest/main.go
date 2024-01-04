package main

import (
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/eliofery/golang-angular/internal/route"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/eliofery/golang-angular/pkg/database/postgres"
	"github.com/gofiber/fiber/v3"
)

func main() {
	conf := config.MustInit(godotenv.New())
	db := database.MustConnect(postgres.New(conf))

	rest := core.New(conf, db)
	rest.SetOptions(fiber.Config{
		ErrorHandler: middleware.NotFound,
	})
	rest.UseMiddlewares(
		middleware.Cors(conf),
	)
	rest.UseRoutes(
		route.NewAuth(),
		route.NewUser(),
	)
	rest.MustRun()
}
