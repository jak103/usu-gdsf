package db

import (
	"errors"
	"os"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

const (
	MOCK      = "mock"
	FIRESTORE = "firestore"
	MONGO     = "mongo"
)

var connection Database

type Database interface {
	GetAllGames() ([]models.Game, error)

	Disconnect() error
	Connect() error
}

// IdExists Helper function to check if given ID already exists in the database
func IdExists(id uint64) (bool, error) {
	_, exists, err := FindById(id)
	return exists, err
}

// FindById Find a game with the given ID in the database
func FindById(id uint64) (models.Game, bool, error) {
	games, err := connection.GetAllGames()
	for _, game := range games {
		if game.Id == id {
			return game, true, err
		}
	}
	return models.Game{}, false, err
}

func NewDatabaseFromEnv() (Database, error) {
	if connection == nil {
		runningEnv, wasSet := os.LookupEnv("RUN_ENV")

		if !wasSet || len(runningEnv) == 0 {
			log.Error("Environment variable RUN_ENV was not set correctly")
			return nil, errors.New("RUN_ENV not set")
		}

		switch runningEnv {
		case MOCK:
			connection = &Mock{}
		case FIRESTORE:
			connection = &Firestore{}
		case MONGO:
			connection = &Mongo{}

		default:
			log.Error("Unknown RUN_ENV set %v", runningEnv)
			return nil, errors.New("unknown RUN_ENV")
		}
	}

	return connection, nil
}
