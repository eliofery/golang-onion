package model

//const RoleTableName = "roles"

const (
	AdminRole = iota + 1
	ModeratorRole
	EditorRole
	UserRole
	GuestRole
)

type Role struct {
	ID          int          `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}
