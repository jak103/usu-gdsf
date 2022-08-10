package models

import (
	"github.com/google/uuid"
)

type Role int64

const (
	Admin     Role = 0
	Publisher      = 1
	Basic          = 2
)

type User struct {
	ID          uuid.UUID `                                      firestore:"id,omitempty"`
	Username    string    `json:"username"    form:"username"    firestore:"username,omitempty"`
	Password    string    `json:"password"    form:"password"    firestore:"password,omitempty"`
	Displayname string    `json:"displayname" form:"displayname" firestore:"displayname,omitempty"`
	Role        Role      `json:"role"        form:"role"        firestore:"role,omitempty"`
}

func (user *User) SetUUID() {
	user.ID = uuid.New()
}
