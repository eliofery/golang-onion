package service

import (
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/gofiber/fiber/v3/log"
)

type UserService interface {
	GetUser()
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	log.Info("инициализация сервиса пользователей")
	return &userService{dao: dao}
}

func (u *userService) GetUser() {}
