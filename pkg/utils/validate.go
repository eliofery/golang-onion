package utils

import (
	"errors"
	"github.com/eliofery/golang-angular/internal/validation"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/gofiber/fiber/v3/log"
)

var (
	ErrLangNotSupported = errors.New("язык не поддерживается")
)

const (
	langDefault = "ru"
)

// Validate валидация данных
type Validate interface {
	Validation(data any, langOptions ...string) []error
	RegisterValidations(customValidates ...validation.CustomValidate)
}

type validate struct {
	Validator *validator.Validate
}

func NewValidate(v *validator.Validate) Validate {
	log.Info("инициализация валидации")
	return &validate{Validator: v}
}

// Validation валидация входных данных
func (v *validate) Validation(data any, langOptions ...string) []error {
	var (
		validatorErr validator.ValidationErrors
		errMessages  []error
	)

	langString := langDefault
	if len(langOptions) > 0 {
		langString = langOptions[0]
	}

	lang := v.setLang(langString)
	if err := v.Validator.Struct(data); err != nil && errors.As(err, &validatorErr) {
		for _, validateErr := range validatorErr {
			errMessage := errors.New(validateErr.Translate(lang))
			errMessages = append(errMessages, errMessage)
		}
	}

	return errMessages
}

// RegisterValidations регистрация пользовательской валидации
func (v *validate) RegisterValidations(customValidates ...validation.CustomValidate) {
	op := "utils.Validate.Register"

	for _, cv := range customValidates {
		err := v.Validator.RegisterValidation(cv.Name, cv.Func)
		if err != nil {
			log.Warnf("%s: %s", op, err)
		}
	}
}

// setLang перевод ошибок валидации
func (v *validate) setLang(lang string) ut.Translator {
	op := "utils.Validate.setLang"

	// TODO: сделать настройку языков более гибкой
	enLang := en.New()
	ruLang := ru.New()
	uni := ut.New(enLang, ruLang)

	trans, ok := uni.GetTranslator(lang)
	if !ok {
		log.Warnf("%s: %s", op, ErrLangNotSupported)
		trans, _ = uni.GetTranslator(langDefault)
	}

	if err := ru_translations.RegisterDefaultTranslations(v.Validator, trans); err != nil {
		log.Warnf("%s: %s", op, err)
	}

	return trans
}
