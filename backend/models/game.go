package models

import (
	"github.com/google/uuid"
)

type Game struct {
	ID               uuid.UUID
	Title            string    `json:"title" form:"title"`
	Description      string    `json:"description" form:"description"`
	UserID           uuid.UUID `json:"userid" form:"userid"`
	VersionNumber    string    `json:"version" form:"version"`
	PublishTimestamp string    `json:"timestamp" form:"timestamp"`
	Tags             []string  `json:"tags" form:"tags"`
}

// todo, create and store unqiue uuids
func (game *Game) SetUUID() {
	game.ID = uuid.New()
}
