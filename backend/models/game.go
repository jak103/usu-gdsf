package models

import (
	"github.com/google/uuid"
)

type ID uuid.UUID
type UserID uuid.UUID

type Game struct {
	ID           uuid.UUID
	Name         string
	Description  string
	UserID       uuid.UUID
	CreationDate string
	Version      string
}

// todo, create and store unqiue uuids
