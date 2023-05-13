package models

import "github.com/google/uuid"

// all the structs used with redis database for cacheing
type Session struct {
	UserId  uuid.UUID `json:"userid"`
	Valid   bool      `json:"valid"`
	IsAdmin bool      `json:"isadmin"`
	Branch  string    `json:"branch"`
}
