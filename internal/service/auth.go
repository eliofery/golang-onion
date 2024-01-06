package service

import (
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/repository"
)

// AuthService содержит бизнес логику авторизации пользователя
type AuthService interface {
	SignUp(user model.User)
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao: dao}
}

// SignUp регистрация пользователя
func (u *authService) SignUp(user model.User) {

}
