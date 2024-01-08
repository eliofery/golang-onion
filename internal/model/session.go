package model

//const SessionTableName = "sessions"

type Sessions struct {
	ID    int    `json:"id,omitempty"`
	Token string `json:"token,omitempty"`
}
