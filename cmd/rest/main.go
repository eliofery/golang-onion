package main

import (
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"github.com/eliofery/golang-angular/pkg/core"
	"github.com/eliofery/golang-angular/pkg/database"
	"github.com/eliofery/golang-angular/pkg/database/postgres"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	conf := config.MustInit(godotenv.New())
	db := database.MustConnect(postgres.New(conf))

	rest := core.New(conf, db)
	rest.UseMiddlewares(
		cors.New(cors.Config{
			AllowOrigins: fmt.Sprintf(
				"%s://%s:%s",
				conf.Get("SERVER_PROTOCOL"), conf.Get("SERVER_URL"), conf.Get("SERVER_PORT"),
			),
			AllowMethods:     "GET, POST, PUT, DELETE",
			AllowHeaders:     "Origin, Content-Type, Accept",
			ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin",
			AllowCredentials: true,
		}),
	)
	rest.UseRoutes()
	rest.MustRun()
}
