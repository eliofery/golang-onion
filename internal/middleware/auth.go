package middleware

import (
	"errors"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"strconv"
)

type key string

const IssuerKey key = "issuer"

var (
	ErrNotAllowed = errors.New("не допустимое действие")
)

// SetUserIdFromToken добавление ID авторизованного пользователя в контекст
func SetUserIdFromToken(dao repository.DAO, tokenManager utils.TokenManager) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		ctx.Locals(IssuerKey, func(c fiber.Ctx) (int, error) {
			message := ErrNotAllowed

			cookieToken := c.Cookies(utils.CookieTokenName)
			authToken := c.Get("Authorization")

			var tokenString string
			if cookieToken != "" {
				tokenString = cookieToken
			} else if authToken != "" {
				tokenString = authToken
			}

			if tokenString == "" {
				return 0, message
			}

			issuer, err := tokenManager.VerifyToken(tokenString)
			if err != nil {
				tokenManager.RemoveCookieToken(c)
				if err = dao.NewSessionQuery().DeleteByToken(tokenString); err != nil {
					log.Errorf("не удалось удалить сессионный токен: %s", err)
				}

				return 0, message
			}

			userId, err := strconv.Atoi(issuer)
			if err != nil {
				log.Errorf("не удалось получить идентификатор пользователя: %v", err)
				return 0, message
			}

			return userId, nil
		})

		return ctx.Next()
	}
}

// IsAuth доступ только для авторизованных пользователей
func IsAuth(ctx fiber.Ctx) error {
	cb, ok := ctx.Locals(IssuerKey).(func(cb fiber.Ctx) (int, error))
	if !ok {
		return ErrNotAllowed
	}

	if _, err := cb(ctx); err != nil {
		return ErrNotAllowed
	}

	return ctx.Next()
}

// IsGuest доступ только для гостей
func IsGuest(ctx fiber.Ctx) error {
	cb, ok := ctx.Locals(IssuerKey).(func(cb fiber.Ctx) (int, error))
	if !ok {
		return ctx.Next()
	}

	if _, err := cb(ctx); err != nil {
		return ctx.Next()
	}

	return ErrNotAllowed
}
