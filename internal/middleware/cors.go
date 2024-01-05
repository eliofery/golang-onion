package middleware

import (
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Cors(conf config.Config) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf(
			"%s://%s:%s",
			conf.Get("SERVER_PROTOCOL"), conf.Get("SERVER_URL"), conf.Get("SERVER_PORT"),
		),
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin",
		AllowCredentials: true,
	})
}