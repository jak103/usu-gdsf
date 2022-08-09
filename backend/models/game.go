package models

import (
	"time"
)
type Game struct {
	Id           string `json:"Id"`
	Name         string `json:"Name" validate:"required"`
	Rating       float32 `json:"Rating" `
	TimesPlayed  int `json:"TimesPlayed" validate:"required, number"`
	ImagePath    string `json:"ImagePath" validate:"required, file"`
	Description  string `json:"Description"`
	Developer    string `json:"Developer" validate:"required, min=5, max=20"`
	CreationDate time.Time `json:"CreationDate"`
	Version      string `json:"Version"`
	Tags         []string `json:"Tags"`
	Downloads    int64 `json:"Downloads" validate:"number"` // firestore doesn't support uint64
	DownloadLink string `json:"DownloadLink" validate:"required, uri"`
}
