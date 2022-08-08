package models

import "time"

type Download struct {
	Id           string
	UserId       string
	GameId       string
	CreationDate time.Time
}