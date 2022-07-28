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
	AddGame(models.Game) (string, error)
	GetGameByID(string) (models.Game, error)
	GetGameID(models.Game) (string, error)
	Disconnect() error
	Connect() error
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
