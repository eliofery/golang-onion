package service

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/bcrypt"
)

// AuthService содержит бизнес логику авторизации пользователя
type AuthService interface {
	SignUp(user dto.UserCreate) (*int, error)
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	log.Info("инициализация AuthService")
	return &authService{dao: dao}
}

// SignUp регистрация пользователя
func (s *authService) SignUp(user dto.UserCreate) (*int, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(passwordHash)
	id, err := s.dao.NewUserQuery().CreateUser(user)
	if err != nil {
		return nil, err
	}

	return id, nil
}
