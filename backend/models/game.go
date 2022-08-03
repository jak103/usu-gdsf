package models

import (
	"time"
)

type Game struct {
	Id           string
	Name         string
	Rating       float32
	TimesPlayed  int
	Description  string
	Developer    string
	CreationDate time.Time
	Version      string
	Tags         []string
	downloads    uint64
}
