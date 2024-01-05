package service

import (
	"github.com/eliofery/golang-angular/internal/repository"
)

type AuthService interface {
	Auth()
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao: dao}
}

func (u *authService) Auth() {}
