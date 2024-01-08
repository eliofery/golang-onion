package validation

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidate пользовательская валидация
type CustomValidate struct {
	Name string
	Func validator.Func
}

// TestValidate пример реализации пользовательской валидации
func TestValidate() CustomValidate {
	return CustomValidate{
		Name: "example",
		Func: func(fl validator.FieldLevel) bool {
			return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
		},
	}
}
