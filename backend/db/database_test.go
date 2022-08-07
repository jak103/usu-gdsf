package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/jak103/usu-gdsf/models"
)

var (
	_db, _ = NewDatabaseFromEnv()

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
)

func TestDatabase_GameID(t *testing.T) {
	// cleanup
	t.Cleanup(cleanup)

	// assign IDs on add
	id0A, _ := _db.AddGame(game0)
	id1A, _ := _db.AddGame(game1)

	game0.Id = id0A
	game1.Id = id1A

	// assigned IDs
	game0A, _ := _db.GetGameByID(id0A)
	game1A, _ := _db.GetGameByID(id1A)
	assert.Equal(t, game0, game0A)
	assert.Equal(t, game1, game1A)
}

func TestDatabase_Tags(t *testing.T) {
	//cleanup
	t.Cleanup(cleanup)

	id0, _ := _db.AddGame(game0)
	id1, _ := _db.AddGame(game1)

	game0.Id = id0
	game1.Id = id1

	res0, _ := _db.GetGamesByTags([]string{"tag0"}, false)
	res1, _ := _db.GetGamesByTags([]string{"tag1"}, false)
	res3, _ := _db.GetGamesByTags([]string{"bad tag"}, false)

	// result size
	assert.Equal(t, 1, len(res0))
	assert.Equal(t, 2, len(res1))
	assert.Equal(t, 0, len(res3))

	// result elements
	assert.Contains(t, res0, game0)
	assert.Contains(t, res1, game1)
	assert.Contains(t, res1, game1)
}
// func TestRemoveGame(t *testing.T){
// 	t.Cleanup(cleanup)
// 	id0, _ := _db.AddGame(game0)
// 	id1, _ := _db.AddGame(game1)

// 	game0.Id = id0
// 	game1.Id = id1

// 	res0, _ := _db.RemoveGame(id0)
// 	assert.Equal(t, res, game0A)

// }
func cleanup() {
	_db.RemoveGame(game0)
	_db.RemoveGame(game1)
}
