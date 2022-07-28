package models

import "time"

type User struct {
	Id              uint64
	Email           string
	PasswordHash    string
	FirstName       string
	LastName        string
	DateOfBirth     time.Time
	JoinedTimestamp time.Time
}
