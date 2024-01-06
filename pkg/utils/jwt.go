package utils

import (
	"fmt"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

const (
	CookieTokenName    = "jwt"
	ExpiresTimeDefault = "3600"
)

var (
	ErrJwtSecretEmpty   = fmt.Errorf("секретный ключ токена не может быть пустым")
	ErrJwtExpires       = fmt.Errorf("не удалось получить время истечения токена")
	ErrJwtCreated       = fmt.Errorf("не удалось создать токен")
	ErrJwtSigningMethod = fmt.Errorf("неожиданный метод подписи токена")
	ErrJwtNotValid      = fmt.Errorf("не верный токен")
)

type Jwt struct {
	conf config.Config
}

func NewJwt(conf config.Config) Jwt {
	return Jwt{conf: conf}
}

// GenerateToken создание токена
func (j *Jwt) GenerateToken(userId int) (string, error) {
	op := "utils.jwt.GenerateToken"

	if j.conf.Get("JWT_SECRET") == "" {
		log.Error(fmt.Errorf("%s: %w", op, ErrJwtSecretEmpty))

		return "", ErrJwtSecretEmpty
	}

	expiresTime, err := j.GetExpiresTime()
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": strconv.Itoa(userId),
		"exp": time.Now().Add(time.Second * expiresTime).Unix(),
	})

	token, err := claims.SignedString([]byte(j.conf.Get("JWT_SECRET")))
	if err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))

		return "", ErrJwtCreated
	}

	return token, nil
}

// VerifyToken валидация токена
func (j *Jwt) VerifyToken(token string) (string, error) {
	op := "utils.jwt.VerifyToken"

	if j.conf.Get("JWT_SECRET") == "" {
		log.Error(fmt.Errorf("%s: %w", op, ErrJwtSecretEmpty))

		return "", ErrJwtSecretEmpty
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error(fmt.Errorf("%s: %w", op, ErrJwtSigningMethod))

			return nil, ErrJwtSigningMethod
		}

		return []byte(j.conf.Get("JWT_SECRET")), nil
	})
	if err != nil || !parsedToken.Valid {
		log.Error(fmt.Errorf("%s: %w", op, err))

		return "", ErrJwtNotValid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Error(fmt.Errorf("%s: %w", op, ErrJwtNotValid))

		return "", ErrJwtNotValid
	}

	issuer, err := claims.GetIssuer()
	if err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))

		return "", ErrJwtNotValid
	}

	return issuer, nil
}

// GetExpiresTime получить время истечения токена
func (j *Jwt) GetExpiresTime() (time.Duration, error) {
	op := "utils.jwt.GetExpiresTime"

	expiresTimeString := os.Getenv("JWT_EXPIRES")
	if expiresTimeString == "" {
		expiresTimeString = ExpiresTimeDefault
	}

	expiresTime, err := strconv.ParseInt(expiresTimeString, 10, 64)
	if err != nil {
		log.Error(fmt.Errorf("%s: %w", op, err))

		return 0, ErrJwtExpires
	}

	return time.Duration(expiresTime), nil
}

// SetCookieToken сохранить токен в куки
func (j *Jwt) SetCookieToken(ctx fiber.Ctx, token string) {
	op := "utils.jwt.SetCookieToken"

	expiresTime, err := j.GetExpiresTime()
	if err != nil {
		log.Warn(fmt.Errorf("%s: %w", op, err))

		return
	}

	cookie := fiber.Cookie{
		Name:     CookieTokenName,
		Value:    token,
		Expires:  time.Now().Add(time.Second * expiresTime),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
}

// RemoveCookieToken удалить токен из куки
func (j *Jwt) RemoveCookieToken(ctx fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     CookieTokenName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
}
