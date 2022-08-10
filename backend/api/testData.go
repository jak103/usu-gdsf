package api

import (
	"time"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/models"
)

// These are the variables used across all api test files.
var (
	_db, _ = db.NewDatabaseFromEnv()

	dummyGameCount = 11

	game0 = models.Game{
		Name:         "game0",
		Rating:       3.5,
		TimesPlayed:  1,
		ImagePath:    "path/0",
		Description:  "dummy game 0",
		Developer:    "tester",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.0",
		Tags:         []string{"tag0", "tag1"},
		Downloads:    35,
		DownloadLink: "dummy.test",
	}

	game1 = models.Game{
		Name:         "game1",
		Rating:       3.9,
		TimesPlayed:  2,
		ImagePath:    "path/1",
		Description:  "dummy game 1",
		Developer:    "tester",
		CreationDate: time.Date(1900, 1, 2, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.1",
		Tags:         []string{"tag1", "tag2"},
		Downloads:    36,
		DownloadLink: "dummy1.test",
	}

	download0 = models.Download{
		UserId:       "000",
		GameId:       "62dadce12fcdb47399118408",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	download1 = models.Download{
		UserId:       "001",
		GameId:       "62dadce12fcdb47399118408",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
	}
)
