package service

import (
	"errors"
	"github.com/eliofery/golang-onion/internal/dto"
	"github.com/eliofery/golang-onion/internal/middleware"
	"github.com/eliofery/golang-onion/internal/model"
	"github.com/eliofery/golang-onion/internal/repository"
	"github.com/eliofery/golang-onion/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/bcrypt"
)

// AuthService содержит логику авторизации пользователя
type AuthService interface {
	GetUserIdFromToken(ctx fiber.Ctx) (userId *int)
	Register(ctx fiber.Ctx, user dto.UserCreate) (token string, err error)
	Auth(ctx fiber.Ctx, user dto.UserAuth) (token string, err error)
	Logout(ctx fiber.Ctx, userId int) error
}

type authService struct {
	dao          repository.DAO
	tokenManager utils.TokenManager
}

func NewAuthService(dao repository.DAO, jwt utils.TokenManager) AuthService {
	log.Info("инициализация сервиса авторизации")
	return &authService{dao: dao, tokenManager: jwt}
}

// GetUserIdFromToken получение идентификатора пользователя из токена
func (s *authService) GetUserIdFromToken(ctx fiber.Ctx) *int {
	cb, ok := ctx.Locals(middleware.IssuerKey).(func(cb fiber.Ctx) (int, error))
	if !ok {
		return nil
	}

	userId, err := cb(ctx)
	if err != nil {
		return nil
	}

	return &userId
}

// Register регистрация и авторизация пользователя
func (s *authService) Register(ctx fiber.Ctx, user dto.UserCreate) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(passwordHash)
	user.RoleID = model.GuestRole
	userId, err := s.dao.NewUserQuery().Create(user)
	if err != nil {
		return "", err
	}

	token, err := s.tokenManager.GenerateToken(userId)
	if err != nil {
		return "", err
	}

	if err = s.dao.NewSessionQuery().Create(userId, token); err != nil {
		return "", err
	}

	s.tokenManager.SetCookieToken(ctx, token)

	return token, nil
}

// Auth авторизация пользователя
func (s *authService) Auth(ctx fiber.Ctx, user dto.UserAuth) (string, error) {
	findUser, err := s.dao.NewUserQuery().GetByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUser.PasswordHash), []byte(user.Password))
	if err != nil {
		return "", errors.New("не верный логин или пароль")
	}

	token, err := s.tokenManager.GenerateToken(findUser.ID)
	if err != nil {
		return "", err
	}

	if err = s.dao.NewSessionQuery().Create(findUser.ID, token); err != nil {
		return "", err
	}

	s.tokenManager.SetCookieToken(ctx, token)

	return token, nil
}

// Logout выход пользователя из системы
func (s *authService) Logout(ctx fiber.Ctx, userId int) error {
	err := s.dao.NewSessionQuery().DeleteByUserId(userId)
	if err != nil {
		return err
	}

	s.tokenManager.RemoveCookieToken(ctx)

	return nil
}
