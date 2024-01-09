package service

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetById(userId int) (user *model.User, err error)
	Create(user dto.UserCreate) (int, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	log.Info("инициализация сервиса пользователей")
	return &userService{dao: dao}
}

// GetById получить пользователя по id
func (s *userService) GetById(userId int) (*model.User, error) {
	user, err := s.dao.NewUserQuery().GetById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Create создать пользователя
func (s *userService) Create(user dto.UserCreate) (int, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(passwordHash)
	userId, err := s.dao.NewUserQuery().Create(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
