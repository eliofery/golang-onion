package service

import (
	"errors"
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/middleware"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/bcrypt"
)

// AuthService содержит логику авторизации пользователя
type AuthService interface {
	GetUserIdFromToken(ctx fiber.Ctx) (userId int)
	Register(user dto.UserCreate) (userId int, err error)
	RegisterAndAuth(ctx fiber.Ctx, user dto.UserCreate) (token string, err error)
	Auth(ctx fiber.Ctx, user dto.UserAuth) (token string, err error)
	Logout(ctx fiber.Ctx, userId int) error
}

type authService struct {
	dao repository.DAO
	jwt utils.TokenManager
}

func NewAuthService(dao repository.DAO, jwt utils.TokenManager) AuthService {
	log.Info("инициализация сервиса авторизации")
	return &authService{dao: dao, jwt: jwt}
}

// GetUserIdFromToken получение идентификатора пользователя из токена
func (s *authService) GetUserIdFromToken(ctx fiber.Ctx) (userId int) {
	cb, ok := ctx.Locals(middleware.IssuerKey).(func(cb fiber.Ctx) (int, error))
	if !ok {
		return 0
	}

	userId, err := cb(ctx)
	if err != nil {
		return 0
	}

	return userId
}

// Register регистрация пользователя
func (s *authService) Register(user dto.UserCreate) (userId int, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(passwordHash)
	userId, err = s.dao.NewUserQuery().Save(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

// RegisterAndAuth регистрация и авторизация пользователя
func (s *authService) RegisterAndAuth(ctx fiber.Ctx, user dto.UserCreate) (token string, err error) {
	userId, err := s.Register(user)
	if err != nil {
		return "", err
	}

	token, err = s.jwt.GenerateToken(userId)
	if err != nil {
		return "", err
	}

	if err = s.dao.NewSessionQuery().Save(userId, token); err != nil {
		return "", err
	}

	s.jwt.SetCookieToken(ctx, token)

	return token, nil
}

// Auth авторизация пользователя
func (s *authService) Auth(ctx fiber.Ctx, user dto.UserAuth) (token string, err error) {
	findUser, err := s.dao.NewUserQuery().GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.PasswordHash), []byte(user.Password))
	if err != nil {
		return "", errors.New("не верный логин или пароль")
	}

	token, err = s.jwt.GenerateToken(findUser.ID)
	if err != nil {
		return "", err
	}

	if err = s.dao.NewSessionQuery().Save(findUser.ID, token); err != nil {
		return "", err
	}

	s.jwt.SetCookieToken(ctx, token)

	return token, nil
}

// Logout выход пользователя из системы
func (s *authService) Logout(ctx fiber.Ctx, userId int) error {
	err := s.dao.NewSessionQuery().DeleteByUserId(userId)
	if err != nil {
		return err
	}

	s.jwt.RemoveCookieToken(ctx)

	return nil
}
