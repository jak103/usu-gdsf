package models

import "time"

type BlacklistedRefreshToken struct {
	Id         uint64
	Token      string
	UserId     uint64
	Expiration int64
	Timestamp  time.Time
}
