package utils

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/gofiber/fiber/v3/log"
)

var (
	enLang = en.New()
	ruLang = ru.New()
	uni    = ut.New(enLang, ruLang)

	ErrLangNotSupported = errors.New("язык не поддерживается")
)

const (
	langDefault = "ru"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

// Validate валидация данных
type Validate struct {
	*validator.Validate
}

func NewValidate(validate *validator.Validate) *Validate {
	return &Validate{Validate: validate}
}

// Validation валидация входных данных
func (v *Validate) Validation(data any, langOptions ...string) []string {
	var (
		validatorErr validator.ValidationErrors
		errMessages  []string
	)

	langString := langDefault
	if len(langOptions) > 0 {
		langString = langOptions[0]
	}

	lang := v.setLang(langString)
	if err := v.Struct(data); err != nil && errors.As(err, &validatorErr) {
		for _, validateErr := range validatorErr {
			errMessages = append(errMessages, validateErr.Translate(lang))
		}
	}

	return errMessages
}

// setLang перевод ошибок валидации
func (v *Validate) setLang(lang string) ut.Translator {
	op := "utils.Validate.setLang"

	trans, ok := uni.GetTranslator(lang)
	if !ok {
		log.Warnf("%s: %s", op, ErrLangNotSupported)
		trans, _ = uni.GetTranslator(langDefault)
	}

	if err := ru_translations.RegisterDefaultTranslations(v.Validate, trans); err != nil {
		log.Warnf("%s: %s", op, err)
	}

	return trans
}

//func (v *Validate) Validation(data any) []ErrorResponse {
//	var validationErrors []ErrorResponse
//
//	if errs := v.Struct(data); errs != nil {
//		for _, err := range errs.(validator.ValidationErrors) {
//			var elem ErrorResponse
//
//			elem.FailedField = err.Field()
//			elem.Tag = err.Tag()
//			elem.Value = err.Value()
//			elem.Error = true
//
//			validationErrors = append(validationErrors, elem)
//		}
//	}
//
//	return validationErrors
//}
//
//func (v *Validate) ErrMessages(errs []ErrorResponse) []string {
//	if len(errs) == 0 && !errs[0].Error {
//		return nil
//	}
//
//	var errMessages []string
//	for _, err := range errs {
//		errMessages = append(errMessages, fmt.Sprintf(
//			"Поле '%s' | Результат '%v' | Ожидание '%s'",
//			err.FailedField,
//			err.Value,
//			err.Tag,
//		))
//	}
//
//	return errMessages
//}
