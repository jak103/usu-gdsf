package models

import (
	"time"
)

type Game struct {
	Id           string
	Name         string
	Rating       float32
	TimesPlayed  int
	ImagePath    string
	Description  string
	Developer    string
	CreationDate time.Time
	Version      string
	Tags         []string
	Downloads    int64 // firestore doesn't support uint64
	DownloadLink string
}
