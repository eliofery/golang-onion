package service

import (
	"github.com/eliofery/golang-angular/internal/validation"
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/gofiber/fiber/v3/log"
)

// ValidateService содержит логику валидации
type ValidateService interface {
	ValidateData(data any, langOptions ...string) []error
	Register(validate ...validation.CustomValidate) *validateService
}

type validateService struct {
	utils.Validate
}

func NewValidateService(validate utils.Validate) ValidateService {
	log.Info("инициализация пользовательской валидации")
	return &validateService{Validate: validate}
}

// ValidateData валидация данных
func (s *validateService) ValidateData(data any, langOptions ...string) []error {
	return s.Validate.Validation(data, langOptions...)
}

// Register регистрация пользовательской валидации
func (s *validateService) Register(validate ...validation.CustomValidate) *validateService {
	s.Validate.RegisterValidations(validate...)

	return s
}
