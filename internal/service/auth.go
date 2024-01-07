package service

import (
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthService содержит бизнес логику авторизации пользователя
type AuthService interface {
	SignUp(user model.User) (*int, error)
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao: dao}
}

// SignUp регистрация пользователя
func (s *authService) SignUp(user model.User) (*int, error) {
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
