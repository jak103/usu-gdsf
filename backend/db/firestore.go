package db

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"google.golang.org/api/iterator"
)

var _ Database = (*Firestore)(nil)

type Firestore struct {
	client *firestore.Client
	games  *firestore.CollectionRef
}

func (db Firestore) GetGameByID(id uuid.UUID) (*models.Game, error) {
	gameDoc := db.games.Doc(id.String())
	docSnapshot, err := gameDoc.Get(context.Background())

	if err != nil {
		return nil, err
	}

	if docSnapshot == nil {
		return nil, fmt.Errorf("%s: game not found", id)
	}

	var game models.Game
	if err = docSnapshot.DataTo(&game); err != nil {
		return nil, err
	}

	return &game, nil
}

func (db Firestore) GetGamesByTags(tags []string) ([]models.Game, error) {
	games := make([]models.Game, 0)
	gc := db.client.Collection("games")

	documents := gc.Where("tags", "in", tags).Documents(context.Background())

	for {
		docRef, docRefErr := documents.Next()

		if docRefErr == iterator.Done {
			break
		}

		var game models.Game
		docRef.DataTo(&game)

		games = append(games, game)
	}

	return games, nil
}

func (db Firestore) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)
	gc := db.client.Collection("games")

	documents := gc.DocumentRefs(context.Background())
	for {
		docRef, docRefErr := documents.Next()

		if docRefErr == iterator.Done {
			break
		}

		var game models.Game

		if docSnapshot, _ := docRef.Get(context.Background()); docSnapshot != nil {
			_ = docSnapshot.DataTo(&game)
		}

		games = append(games, game)
	}

	return games, nil
}

// func (d *Firestore) GetGamesByTags(tags []string) ([]models.Game, error) {
// 	panic("not implemented") // TODO: Implement
// }

func (db *Firestore) GetGamesByPublishDate(startRange string, endRange string) ([]models.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) CreateGame(newGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) DeleteGame(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) UpdateGame(updatedGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

// Users
func (db *Firestore) GetAllUsers() ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) GetUserByID(id uuid.UUID) (*models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) GetUserByUserName(userName string) (*models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) GetUsersByRole(role int64) ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) CreateUser(newUser models.User) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) DeleteUser(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) UpdateUser(updatedUser models.User) error {
	panic("not implemented") // TODO: Implement
}

// Ratings
func (db *Firestore) GetRatingByID(id uuid.UUID) (*models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) GetRatingsByGame(gameID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) GetRatingsByUser(userID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) CreateRating(newRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) DeleteRating(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) DeleteRatingsByGame(gameID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) UpdateRating(updatedRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
}

// Disconnect disconnects from the remote database
func (db *Firestore) Disconnect() error {
	// Close the client connection if it is open
	if db.client != nil {
		if err := db.client.Close(); err != nil {
			log.WithError(err).Error("Failed to disconnect firestore")
			return err
		}
	}

	return nil
}

// Connect allows the user to connect to the database
func (db *Firestore) Connect() error {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("FIRESTORE_PROJECT_ID")

	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		log.WithError(err).Error("Failed to get firestore client")
		return err
	}

	// Etablish Database Collection object
	db.client = client

	return nil
}
