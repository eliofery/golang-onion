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

// TokenManager jwt токен
type TokenManager interface {
	GenerateToken(userId int) (token string, err error)
	VerifyToken(token string) (issuer string, err error)
	GetExpiresTime() (time.Duration, error)
	SetCookieToken(ctx fiber.Ctx, token string)
	RemoveCookieToken(ctx fiber.Ctx)
}

type tokenManager struct {
	conf config.Config
}

func NewTokenManager(conf config.Config) TokenManager {
	log.Info("инициализация tokenManager")
	return &tokenManager{conf: conf}
}

// GenerateToken создание токена
func (t *tokenManager) GenerateToken(userId int) (token string, err error) {
	op := "utils.jwt.GenerateToken"

	if t.conf.Get("JWT_SECRET") == "" {
		log.Errorf("%s: %s", op, ErrJwtSecretEmpty)
		return "", ErrJwtSecretEmpty
	}

	expiresTime, err := t.GetExpiresTime()
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": strconv.Itoa(userId),
		"exp": time.Now().Add(time.Second * expiresTime).Unix(),
	})

	token, err = claims.SignedString([]byte(t.conf.Get("JWT_SECRET")))
	if err != nil {
		log.Errorf("%s: %s", op, err)
		return "", ErrJwtCreated
	}

	return token, nil
}

// VerifyToken валидация токена
func (t *tokenManager) VerifyToken(token string) (issuer string, err error) {
	op := "utils.jwt.VerifyToken"

	if t.conf.Get("JWT_SECRET") == "" {
		log.Errorf("%s: %s", op, ErrJwtSecretEmpty)
		return "", ErrJwtSecretEmpty
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Errorf("%s: %s", op, ErrJwtSigningMethod)
			return nil, ErrJwtSigningMethod
		}

		return []byte(t.conf.Get("JWT_SECRET")), nil
	})
	if err != nil || !parsedToken.Valid {
		log.Errorf("%s: %s", op, err)
		return "", ErrJwtNotValid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Errorf("%s: %s", op, ErrJwtNotValid)
		return "", ErrJwtNotValid
	}

	issuer, err = claims.GetIssuer()
	if err != nil {
		log.Errorf("%s: %s", op, err)
		return "", ErrJwtNotValid
	}

	return issuer, nil
}

// GetExpiresTime получить время истечения токена
func (t *tokenManager) GetExpiresTime() (time.Duration, error) {
	op := "utils.jwt.GetExpiresTime"

	expiresTimeString := os.Getenv("JWT_EXPIRES")
	if expiresTimeString == "" {
		expiresTimeString = ExpiresTimeDefault
	}

	expiresTime, err := strconv.ParseInt(expiresTimeString, 10, 64)
	if err != nil {
		log.Errorf("%s: %s", op, err)
		return 0, ErrJwtExpires
	}

	return time.Duration(expiresTime), nil
}

// SetCookieToken сохранить токен в куки
func (t *tokenManager) SetCookieToken(ctx fiber.Ctx, token string) {
	op := "utils.jwt.SetCookieToken"

	expiresTime, err := t.GetExpiresTime()
	if err != nil {
		log.Warnf("%s: %s", op, err)
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
func (t *tokenManager) RemoveCookieToken(ctx fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     CookieTokenName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
}
