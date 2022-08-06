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
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Displayname string `json:"displayname" form:"displayname" `
	Role        Role   `json:"role" form:"role"`
}

func (user *User) SetUUID() {
	user.ID = uuid.New()
}
