package service

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/bcrypt"
)

// AuthService содержит логику авторизации пользователя
type AuthService interface {
	Register(user dto.UserCreate) (id int, err error)
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
