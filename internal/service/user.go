package service

import (
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/gofiber/fiber/v3/log"
)

type UserService interface {
	GetUser(userId int) (user *model.User, err error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	log.Info("инициализация сервиса пользователей")
	return &userService{dao: dao}
}

func (u *userService) GetUser(userId int) (*model.User, error) {
	user, err := u.dao.NewUserQuery().GetUserById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}
