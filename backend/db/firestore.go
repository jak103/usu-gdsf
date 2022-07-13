package db

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"google.golang.org/api/iterator"
)

var _ Database = (*Firestore)(nil)

type Firestore struct {
	client *firestore.Client
}

<<<<<<< HEAD:backend/db/FirestoreDB.go
func (db *firestoreDB) GetAllGameRecords() (*[]models.GameRecord, error) {
	games := make([]models.GameRecord, 0)
=======
func (db Firestore) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)
	gc := db.client.Collection("games")
>>>>>>> cfdfbb2712e991bd0e65505a5f7bad6a8f66f521:backend/db/firestore.go

	documents := gc.DocumentRefs(context.Background())
	for {
		docRef, docRefErr := documents.Next()

		if docRefErr == iterator.Done {
			break
		}

<<<<<<< HEAD:backend/db/FirestoreDB.go
		var game models.GameRecord
=======
		var game models.Game
>>>>>>> cfdfbb2712e991bd0e65505a5f7bad6a8f66f521:backend/db/firestore.go

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

<<<<<<< HEAD:backend/db/FirestoreDB.go
func init() {
	registerDB(&DB{
		Name:          "FIRESTORE",
		Description:   "Production Firestore connection",
		StoreDatabase: new(firestoreDB),
	})
=======
	return nil
>>>>>>> cfdfbb2712e991bd0e65505a5f7bad6a8f66f521:backend/db/firestore.go
}
