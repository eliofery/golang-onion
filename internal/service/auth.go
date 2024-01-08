package service

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/bcrypt"
)

// AuthService содержит логику авторизации пользователя
type AuthService interface {
	Register(user dto.UserCreate) (id int, err error)
	RegisterAndAuth(ctx fiber.Ctx, user dto.UserCreate) (token string, err error)
}

type authService struct {
	dao repository.DAO
	jwt utils.TokenManager
}

func NewAuthService(dao repository.DAO, jwt utils.TokenManager) AuthService {
	log.Info("инициализация сервиса авторизации")
	return &authService{dao: dao, jwt: jwt}
}

// Register регистрация пользователя
func (s *authService) Register(user dto.UserCreate) (id int, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(passwordHash)
	id, err = s.dao.NewUserQuery().Save(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// RegisterAndAuth регистрация и авторизация пользователя
func (s *authService) RegisterAndAuth(ctx fiber.Ctx, user dto.UserCreate) (token string, err error) {
	id, err := s.Register(user)
	if err != nil {
		return "", err
	}

	token, err = s.jwt.GenerateToken(id)
	if err != nil {
		return "", err
	}

	s.jwt.SetCookieToken(ctx, token)

	return token, nil
}
