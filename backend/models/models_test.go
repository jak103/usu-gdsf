package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Testing for downloads.go
func TestDownloadsModel(t *testing.T) {
	d1 := Download{
		Id:           "Su22#000456",
		UserId:       "Admin#0002",
		GameId:       "1",
		CreationDate: time.Date(2022, time.August, 8, 19, 2, 15, 10, time.Local),
	}

	// Checks a fully built model
	assert.Equal(t, d1.Id, "Su22#000456")
	assert.Equal(t, d1.UserId, "Admin#0002")
	assert.Equal(t, d1.GameId, "1")
	assert.False(t, time.Time.IsZero(d1.CreationDate))

	d2 := Download{
		Id:     "Su22#000457",
		UserId: "User#8417",
	}

	// Checks a partially built model
	assert.Equal(t, d2.Id, "Su22#000457")
	assert.Equal(t, d2.UserId, "User#8417")
	assert.Equal(t, d2.GameId, "")
	assert.True(t, time.Time.IsZero(d2.CreationDate))
}
