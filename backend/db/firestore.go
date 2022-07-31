package db

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"google.golang.org/api/iterator"
)

var _ Database = (*Firestore)(nil)

type Firestore struct {
	client *firestore.Client
	games  *firestore.CollectionRef
}

func (db Firestore) GetGameByID(id string) (*models.Game, error) {
	gameDoc := db.games.Doc(id)
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
