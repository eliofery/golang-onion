package service

import (
	"github.com/eliofery/golang-onion/internal/dto"
	"github.com/eliofery/golang-onion/internal/repository"
	"github.com/eliofery/golang-onion/pkg/config"
	"github.com/gofiber/fiber/v3/log"
	"math"
	"strconv"
)

type PermissionService interface {
	GetAll(page int) (permissions *dto.PermissionAll, err error)
}

type permissionService struct {
	dao  repository.DAO
	conf config.Config
}

func NewPermissionService(dao repository.DAO, conf config.Config) PermissionService {
	log.Info("инициализация сервиса разрешений")
	return &permissionService{dao: dao, conf: conf}
}

// GetAll получить разрешения
func (s *permissionService) GetAll(page int) (*dto.PermissionAll, error) {
	const defaultLimit = 10

	limit, err := strconv.Atoi(s.conf.Get("PAGINATION_LIMIT"))
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset := (page - 1) * limit

	permissions, err := s.dao.NewPermissionQuery().GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.dao.NewPermissionQuery().GetTotalCount()
	if err != nil {
		return nil, err
	}

	var result dto.PermissionAll
	result.Permissions = permissions
	result.Meta.Total = total
	result.Meta.Page = page
	result.Meta.LastPage = math.Ceil(float64(total) / float64(limit))

	return &result, err
}
