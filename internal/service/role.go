package service

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/gofiber/fiber/v3/log"
	"math"
	"strconv"
)

type RoleService interface {
	GetAll(page int) (roles *dto.RoleAll, err error)
}

type roleService struct {
	dao  repository.DAO
	conf config.Config
}

func NewRoleService(dao repository.DAO, conf config.Config) RoleService {
	log.Info("инициализация сервиса ролей")
	return &roleService{dao: dao, conf: conf}
}

func (s *roleService) GetAll(page int) (*dto.RoleAll, error) {
	const defaultLimit = 10

	limit, err := strconv.Atoi(s.conf.Get("PAGINATION_LIMIT"))
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset := (page - 1) * limit

	roles, err := s.dao.NewRoleQuery().GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.dao.NewRoleQuery().GetTotalCount()
	if err != nil {
		return nil, err
	}

	var result dto.RoleAll
	result.Roles = roles
	result.Meta.Total = total
	result.Meta.Page = page
	result.Meta.LastPage = math.Ceil(float64(total) / float64(limit))

	return &result, err
}
