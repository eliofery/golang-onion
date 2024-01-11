package dto

import "github.com/eliofery/golang-angular/internal/model"

type Permission struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
}

type PermissionAll struct {
	Permissions []model.Permission `json:"permissions"`
	Meta        struct {
		Total    int     `json:"total"`
		Page     int     `json:"page"`
		LastPage float64 `json:"last_page"`
	} `json:"meta"`
}
