package dto

import "github.com/eliofery/golang-angular/internal/model"

type userRequiredFields struct {
	ID              int    `json:"id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password,omitempty" validate:"required"`
	PasswordConfirm string `json:"password_confirm,omitempty" validate:"required,eqfield=Password"`
}

type userNotRequiredFields struct {
	ID              int    `json:"id,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Email           string `json:"email" validate:"email"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"password_confirm,omitempty" validate:"eqfield=Password"`
}

type UserCreate struct {
	userRequiredFields
}

type UserUpdate struct {
	userNotRequiredFields
}

type UserAuth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UserAll struct {
	Users []model.User `json:"users"`
	Meta  struct {
		Total    int     `json:"total"`
		Page     int     `json:"page"`
		LastPage float64 `json:"last_page"`
	} `json:"meta"`
}
