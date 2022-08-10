package models

import "time"

type User struct {
	Email           string
	PasswordHash    string
	FirstName       string
	LastName        string
	DateOfBirth     time.Time
	JoinedTimestamp time.Time
	Role            string
}
