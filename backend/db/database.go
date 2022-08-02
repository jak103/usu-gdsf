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
	GetGameByID(id string) (*models.Game, error)

	Disconnect() error
	Connect() error
}

func NewDatabaseFromEnv() (Database, error) {
	if connection == nil {
		dbType, wasSet := os.LookupEnv("DB_TYPE")

		if !wasSet || len(dbType) == 0 {
			log.Error("Environment variable DB_TYPE was not set correctly")
			return nil, errors.New("DB_TYPE not set")
		}

		switch dbType {
		case MOCK:
			connection = &Mock{}
		case FIRESTORE:
			connection = &Firestore{}
		case MONGO:
			connection = &Mongo{}

		default:
			log.Error("Unknown DB_TYPE set %v", dbType)
			return nil, errors.New("unknown DB_TYPE")
		}
	}

	return connection, nil
}
