package models

import (
	"github.com/google/uuid"
)

type Game struct {
	GameId       uuid.UUID
	Title        string
	Description  string
	UserID       uuid.UUID
	VersionNumber string
}

// todo, create and store unqiue uuids
