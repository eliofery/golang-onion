package service

import (
	"github.com/eliofery/golang-angular/internal/repository"
)

type UserService interface {
	GetUser()
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) GetUser() {
	err := u.dao.NewUserQuery().GetUser()
	if err != nil {
		return
	}
}
