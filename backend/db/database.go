package db

import (
	"errors"

	"github.com/jak103/usu-gdsf/config"
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
	RemoveGame(models.Game) error
	GetGameByID(string) (models.Game, error)
	GetGamesByTag(string) ([]models.Game, error)
	GetGameID(models.Game) (string, error)
	Disconnect() error
	Connect() error
}

func NewDatabaseFromEnv() (Database, error) {
	if connection == nil {
		dbType := config.DbType

		switch dbType {
		case MOCK:
			connection = &Mock{}
		case FIRESTORE:
			connection = &Firestore{}
		case MONGO:
			connection = &Mongo{}

		default:
			log.Error("Unknown DB_TYPE %v", dbType)
			return nil, errors.New("unknown RUN_ENV")
		}

		err := connection.Connect()
		if err != nil {
			log.WithError(err).Error("Could not connect to database")
			return nil, err
		}
	}
	return connection, nil
}
