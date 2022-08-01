package models

import (
	"github.com/google/uuid"
)

type GameRating struct {
	ID                 uuid.UUID
	RatingValue        string
	RatingDescription  string
	RatingTimestamp    string
	GameId             uuid.UUID
	UserID             uuid.UUID
}
