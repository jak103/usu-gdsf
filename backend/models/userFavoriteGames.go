package models

import "time"

type UserFavoriteGame struct {
	Id        uint64
	UserId    uint64
	GameId    uint64
	Timestamp time.Time
}
