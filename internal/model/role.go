package model

//const RoleTableName = "roles"

type Role struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
