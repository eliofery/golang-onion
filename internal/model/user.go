package model

//const UserTableName = "users"

type User struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
	RoleID       int    `json:"role_id,omitempty"`
	Role         Role   `json:"role,omitempty"`
}
