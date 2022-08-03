package models

import (
	"github.com/google/uuid"
)


type Role int64

const (
	Admin Role  = 0
	Publisher   = 1

)

type User struct {

	ID       uuid.UUID
	Username     string
	EmailAddress string
	Password     string
	Displayname  string
	Role         Role
}



