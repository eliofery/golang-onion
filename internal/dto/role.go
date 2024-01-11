package dto

import "github.com/eliofery/golang-angular/internal/model"

type Role struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name,omitempty" validate:"required"`
	Permissions []Permission `json:"permissions,omitempty"`
}

type RolePermission struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required"`
	Permissions []int  `json:"permissions,omitempty" validate:"required"`
}

type RoleAll struct {
	Roles []model.Role `json:"roles"`
	Meta  struct {
		Total    int     `json:"total"`
		Page     int     `json:"page"`
		LastPage float64 `json:"last_page"`
	} `json:"meta"`
}
