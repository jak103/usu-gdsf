package models

import (
	"github.com/google/uuid"
)

type GameTag struct {
	TagId              uuid.UUID
	Type               string
	GameId             uuid.UUID
	GameRatingId       uuid.UUID
	UserID             uuid.UUID
}
