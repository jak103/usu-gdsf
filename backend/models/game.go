package models

import (
	"time"
)

type Game struct {
	Id            string    `json:"Id"`
	Name          string    `json:"Name"`
	Rating        float32   `json:"Rating"`
	TimesPlayed   int       `json:"TimesPlayed"`
	ImagePath     string    `json:"ImagePath"`
	Description   string    `json:"Description"`
	Developer     string    `json:"Developer"`
	CreationDate  time.Time `json:"CreationDate"`
	Version       string    `json:"Version"`
	Tags          []string  `json:"Tags"`
	Downloads     int64     `json:"Downloads"` // firestore doesn't support uint64
	DownloadLink  string    `json:"DownloadLink"`
	reviewIds     []string  `json:"reviewIds"`
	averageReview float64   `json:"averageReview"`
}
