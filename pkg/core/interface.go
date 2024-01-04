package core

import "github.com/gofiber/fiber/v3"

type Route interface {
	Setup(*fiber.App)
}
