package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID
	Username     string
	Password     string
	UserID       uuid.UUID
	Displayname  string
	RolesId      string
}
