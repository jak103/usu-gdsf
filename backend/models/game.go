package models

import (
	"github.com/google/uuid"
)

type Game struct {
	ID               uuid.UUID `                                      firestore:"id,omitempty"`
	Title            string    `json:"title"       form:"title"       firestore:"title,omitempty"`
	Description      string    `json:"description" form:"description" firestore:"description,omitempty"`
	UserID           uuid.UUID `json:"userid"      form:"userid"      firestore:"userid,omitempty"`
	VersionNumber    string    `json:"version"     form:"version"     firestore:"version,omitempty"`
	PublishTimestamp string    `json:"timestamp"   form:"timestamp"   firestore:"timestamp,omitempty"`
	Tags             []string  `json:"tags"        form:"tags"        firestore:"tags,omitempty"`
}

// todo, create and store unqiue uuids
func (game *Game) SetUUID() {
	game.ID = uuid.New()
}
