package service

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strconv"
)

type UserService interface {
	GetById(userId int) (user *model.User, err error)
	Create(user dto.UserCreate) (int, error)
	GetAll(page int) (*dto.UserAll, error)
	Update(user dto.UserUpdate) (*model.User, error)
}

type userService struct {
	dao  repository.DAO
	conf config.Config
}

func NewUserService(dao repository.DAO, conf config.Config) UserService {
	log.Info("инициализация сервиса пользователей")
	return &userService{dao: dao, conf: conf}
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

// GetAll получить всех пользователей
func (s *userService) GetAll(page int) (*dto.UserAll, error) {
	const defaultLimit = 10

	limit, err := strconv.Atoi(s.conf.Get("PAGINATION_LIMIT"))
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset := (page - 1) * limit

	users, err := s.dao.NewUserQuery().GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.dao.NewUserQuery().GetTotalCount()
	if err != nil {
		return nil, err
	}

	var result dto.UserAll
	result.Users = users
	result.Meta.Total = total
	result.Meta.Page = page
	result.Meta.LastPage = math.Ceil(float64(total) / float64(limit))

	return &result, err
}

// Update обновить пользователя
func (s *userService) Update(user dto.UserUpdate) (*model.User, error) {
	if user.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(passwordHash)
	}

	updateUser, err := s.dao.NewUserQuery().Update(user)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}
