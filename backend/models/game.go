package models

import (
	"github.com/google/uuid"
)

type Game struct {
	ID      uuid.UUID
	Title        string
	Description  string
	UserID       uuid.UUID
	VersionNumber string
}

// todo, create and store unqiue uuids
