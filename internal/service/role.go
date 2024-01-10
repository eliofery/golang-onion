package service

import (
	"github.com/eliofery/golang-angular/internal/dto"
	"github.com/eliofery/golang-angular/internal/model"
	"github.com/eliofery/golang-angular/internal/repository"
	"github.com/eliofery/golang-angular/pkg/config"
	"github.com/gofiber/fiber/v3/log"
	"math"
	"strconv"
)

type RoleService interface {
	GetAll(page int) (roles *dto.RoleAll, err error)
	GetById(roleId int) (role *model.Role, err error)
	Update(role dto.Role) (updateRole *model.Role, err error)
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

// GetById получить роли по id
func (s *roleService) GetById(roleId int) (*model.Role, error) {
	role, err := s.dao.NewRoleQuery().GetById(roleId)
	if err != nil {
		return role, err
	}

	return role, nil
}

// Update обновить роль
func (s *roleService) Update(role dto.Role) (*model.Role, error) {
	updateRole, err := s.dao.NewRoleQuery().Update(role)
	if err != nil {
		return nil, err
	}

	return updateRole, nil
}
