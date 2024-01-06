package model

import (
	"github.com/eliofery/golang-angular/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3/log"
)

// Validate дополняет стандартную валидацию
type Validate struct {
	*utils.Validate
}

func NewValidate(validate *utils.Validate) Validate {
	var v Validate
	v.Validate = validate
	v.registerValidations()

	return v
}

// registerValidations регистрация пользовательских валидаций
func (v *Validate) registerValidations() {
	v.testValidate()
}

// testValidate пример реализации пользовательской валидации
func (v *Validate) testValidate() {
	op := "model.Validate.testValidate"

	err := v.Validate.RegisterValidation("test_validate", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})
	if err != nil {
		log.Warnf("%s: %s", op, err)
	}
}
