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

// Testing for game.go
func TestGameModel(t *testing.T) {
	g1 := Game{
		Id:           "1",
		Name:         "Minecraft",
		Rating:       4.5,
		TimesPlayed:  40000000,
		Description:  "Explore some blocks",
		Developer:    "Mojang",
		CreationDate: time.Date(2009, time.May, 1, 0, 0, 0, 0, time.UTC),
		Version:      "1.x",
		Tags:         []string{"Sandbox", "Adventure"},
		Downloads:    160000000,
		DownloadLink: "www.minecraft.com",
	}

	// Checks a fully built model
	assert.Equal(t, g1.Id, "1")
	assert.Equal(t, g1.Name, "Minecraft")
	assert.Equal(t, g1.Rating, float32(4.5))
	assert.Equal(t, g1.TimesPlayed, 40000000)
	assert.Equal(t, g1.Description, "Explore some blocks")
	assert.Equal(t, g1.Developer, "Mojang")
	assert.False(t, time.Time.IsZero(g1.CreationDate))
	assert.Equal(t, g1.Version, "1.x")
	assert.Equal(t, g1.Tags, []string{"Sandbox", "Adventure"})
	assert.Equal(t, g1.Downloads, int64(160000000))
	assert.Equal(t, g1.DownloadLink, "www.minecraft.com")

	g2 := Game{
		Id:           "2",
		TimesPlayed:  1200,
		Description:  "Exciting first person shooter",
		Developer:    "Activision",
		Version:      "1.x",
		Tags:         []string{"Spring2020", "FPS"},
		DownloadLink: "www.activision.com",
	}

	// Checks a partially built model
	assert.Equal(t, g2.Id, "2")
	assert.Equal(t, g2.Name, "")
	assert.Equal(t, g2.Rating, float32(0.0))
	assert.Equal(t, g2.TimesPlayed, 1200)
	assert.Equal(t, g2.Description, "Exciting first person shooter")
	assert.Equal(t, g2.Developer, "Activision")
	assert.True(t, time.Time.IsZero(g2.CreationDate))
	assert.Equal(t, g2.Version, "1.x")
	assert.Equal(t, g2.Tags, []string{"Spring2020", "FPS"})
	assert.Equal(t, g2.Downloads, int64(0))
	assert.Equal(t, g2.DownloadLink, "www.activision.com")
}

func TestUserModel(t *testing.T) {
	u1 := User{
		Email:           "john.doe@gmail.com",
		PasswordHash:    "dkj37%39rhI83)",
		FirstName:       "John",
		LastName:        "Doe",
		DateOfBirth:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local),
		JoinedTimestamp: time.Date(2022, time.August, 8, 4, 12, 15, 1, time.Local),
	}

	// Checks a fully built model
	assert.Equal(t, u1.Email, "john.doe@gmail.com")
	assert.Equal(t, u1.PasswordHash, "dkj37%39rhI83)")
	assert.Equal(t, u1.FirstName, "John")
	assert.Equal(t, u1.LastName, "Doe")
	assert.False(t, time.Time.IsZero(u1.DateOfBirth))
	assert.False(t, time.Time.IsZero(u1.JoinedTimestamp))

	u2 := User{
		Email:       "jane.doe@gmail.com",
		FirstName:   "Jane",
		LastName:    "Doe",
		DateOfBirth: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local),
	}

	// Checks a partially built model
	assert.Equal(t, u2.Email, "jane.doe@gmail.com")
	assert.Equal(t, u2.PasswordHash, "")
	assert.Equal(t, u2.FirstName, "Jane")
	assert.Equal(t, u2.LastName, "Doe")
	assert.False(t, time.Time.IsZero(u2.DateOfBirth))
	assert.True(t, time.Time.IsZero(u2.JoinedTimestamp))
}
