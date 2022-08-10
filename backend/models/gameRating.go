package models

import (
	"github.com/google/uuid"
)

type GameRating struct {
	ID                uuid.UUID `firestore:"id,omitempty"`
	RatingValue       int       `firestore:"ratingvalue,omitempty"`
	RatingDescription string    `firestore:"ratingdescription,omitempty"`
	RatingTimestamp   string    `firestore:"ratingtimestamp,omitempty"`
	GameId            uuid.UUID `firestore:"gameid,omitempty"`
	UserID            uuid.UUID `firestore:"userid,omitempty"`
}

func (gameRating *GameRating) SetUUID() {
	gameRating.ID = uuid.New()
}
