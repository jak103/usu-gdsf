package db

import (
	"testing"
	"time"

	"github.com/jak103/usu-gdsf/models"
	"github.com/stretchr/testify/assert"
)

var (
	game0 = models.Game{
		Name:         "game0",
		Developer:    "tester",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.0",
		Tags:         []string{"tag9", "tag10"},
	}

	game1 = models.Game{
		Name:         "game1",
		Developer:    "tester",
		CreationDate: time.Date(1900, 1, 2, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.1",
		Tags:         []string{"tag11", "tag12"},
	}
)

var (
	game2 = models.Game{
		Name:         "game2",
		Developer:    "tester",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.0",
		Tags:         []string{"tag6", "tag7"},
	}

	game3 = models.Game{
		Name:         "game3",
		Developer:    "tester",
		CreationDate: time.Date(1900, 1, 2, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.1",
		Tags:         []string{"tag7", "tag8"},
	}
)

func TestDatabase_GameID(t *testing.T) {
	_db, _ := NewDatabaseFromEnv()

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

	// cleanup
	_db.RemoveGame(game0)
	_db.RemoveGame(game1)
}

func TestDatabase_Tags(t *testing.T) {
	_db, _ := NewDatabaseFromEnv()
	id0, _ := _db.AddGame(game2)
	id1, _ := _db.AddGame(game3)

	game2.Id = id0
	game3.Id = id1

	res0, _ := _db.GetGamesByTags([]string{"tag6"}, false)
	res1, _ := _db.GetGamesByTags([]string{"tag7"}, false)
	res3, _ := _db.GetGamesByTags([]string{"bad tag"}, false)

	// result size
	assert.Equal(t, 1, len(res0))
	assert.Equal(t, 2, len(res1))
	assert.Equal(t, 0, len(res3))

	// result elements
	assert.Contains(t, res0, game2)
	assert.Contains(t, res1, game2)
	assert.Contains(t, res1, game3)

	// cleanup
	_db.RemoveGame(game2)
	_db.RemoveGame(game3)
}
