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
	GetUserByID(id uuid.UUID)
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
	updateRating(updatedRating models.GameRating) error

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
