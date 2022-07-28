package models

import "time"

type Download struct {
	Id        uint64
	UserId    uint64
	GameId    uint64
	Timestamp time.Time
}