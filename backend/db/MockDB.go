package db

import (
	// "errors"
	// "fmt"

	// "github.com/google/uuid"
	"strconv"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
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
}

// connect allows the user to connect to the database
func (db *mockDB) connect() {
	if len(db.gameRecords) == 0 {
		for i, v := range CreateGamesFromJson() {
			db.gameRecords[strconv.Itoa(i)] = v
		}

		log.Debug("Sample Database initialized.")
	}
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
