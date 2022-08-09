package db

import (
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

const (
	MOCK      = "mock"
	FIRESTORE = "firestore"
	MONGO     = "mongo"
)

var connection Database
var DB Database

type Database interface {
	// Games
	GetAllGames() ([]models.Game, error)
	GetGameByID(id uuid.UUID) (*models.Game, error)
	GetGamesByTags(tags []string) ([]models.Game, error)
	GetGamesByPublishDate(startRange string, endRange string) ([]models.Game, error)
	CreateGame(newGame models.Game) error
	DeleteGame(id uuid.UUID) error
	UpdateGame(updatedGame models.Game) error

	// Users
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetUserByUserName(userName string) (*models.User, error)
	GetUsersByRole(role int64) ([]models.User, error)
	CreateUser(newUser models.User) error
	DeleteUser(id uuid.UUID) error
	UpdateUser(updatedUser models.User) error

	// Ratings
	GetRatingByID(id uuid.UUID) (*models.GameRating, error)
	GetRatingsByGame(gameID uuid.UUID) ([]models.GameRating, error)
	GetRatingsByUser(userID uuid.UUID) ([]models.GameRating, error)
	CreateRating(newRating models.GameRating) error
	DeleteRating(id uuid.UUID) error
	DeleteRatingsByGame(gameID uuid.UUID) error
	UpdateRating(updatedRating models.GameRating) error

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
	DB = connection
	return connection, nil
}
