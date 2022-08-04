package models

import (
	"github.com/google/uuid"
)

type Role int64

const (
	Admin     Role = 0
	Publisher Role = 1
)

type User struct {
	ID          uuid.UUID
	Username    string
	Password    string
	Displayname string
	Role        Role
}

func (user *User) SetUUID() {
	user.ID = uuid.New()
}
