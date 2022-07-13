package db

import (
	// "errors"
	// "fmt"

	// "github.com/google/uuid"
	"github.com/jak103/usu-gdsf/models"
	"github.com/jak103/usu-gdsf/log"
)

// MockDB is an implemenation declaring the unit test db
type mockDB struct {
	gameRecords         map[string]models.GameRecord
}

func (db *mockDB) GetAllGameRecords() (*[]models.GameRecord, error) {
	games := make([]models.GameRecord, 0)

	for _, game := range db.gameRecords {
		games = append(games, game)
	}

	return &games, nil
}

// Disconnect disconnects from the remote database
func (db *mockDB) disconnect() {
	return
}

// connect allows the user to connect to the database
func (db *mockDB) connect() {
	if len(db.gameRecords) == 0 {
		thing := CreateGamesFromJson()

		for _, v := range thing {
			log.Debug("", v)
		}
	}
	return
}

func init() {
	registerDB(&DB{
		Name:          "MOCK",
		Description:   "Mock database connection for Unit Tests",
		StoreDatabase: &mockDB{
			gameRecords:         make(map[string]models.GameRecord),
		},
	})
}
