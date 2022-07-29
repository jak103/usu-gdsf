package db

import (
	"github.com/jak103/usu-gdsf/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	game0 = models.Game{
		Name:         "game0",
		Author:       "tester",
		CreationDate: "01/01/1900",
		Version:      "0.0.0",
		Tags:         nil,
	}

	game1 = models.Game{
		Name:         "game1",
		Author:       "tester",
		CreationDate: "01/02/1900",
		Version:      "0.0.1",
		Tags:         nil,
	}
)

func TestMongo_GameID(t *testing.T) {
	_db, _ := NewDatabaseFromEnv()

	// assign IDs on add
	id0A, _ := _db.AddGame(game0)
	id1A, _ := _db.AddGame(game1)

	// find IDs with game details
	id0F, _ := _db.GetGameID(game0)
	id1F, _ := _db.GetGameID(game1)

	// assigned IDs
	game0A, _ := _db.GetGameByID(id0A)
	game1A, _ := _db.GetGameByID(id1A)
	assert.Equal(t, game0, game0A)
	assert.Equal(t, game1, game1A)

	// found IDs
	game0F, _ := _db.GetGameByID(id0F)
	game1F, _ := _db.GetGameByID(id1F)
	assert.Equal(t, game0, game0F)
	assert.Equal(t, game1, game1F)
}
