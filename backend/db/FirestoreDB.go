package db

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/jak103/uno/model"
	"google.golang.org/api/iterator"
)

type firestoreDB struct {
	client  *firestore.Client
	games   *firestore.CollectionRef
	players *firestore.CollectionRef
}

func (db *firestoreDB) GetAllGames() (*[]model.Game, error) {
	games := make([]model.Game, 0)

	documents := db.games.DocumentRefs(context.Background())
	for {
		docRef, docRefErr := documents.Next()

		if docRefErr == iterator.Done {
			break
		}

		var game model.Game

		if docSnapshot, _ := docRef.Get(context.Background()); docSnapshot != nil {
			_ = docSnapshot.DataTo(&game)
		}

		games = append(games, game)
	}

	return &games, nil
}

// Disconnect disconnects from the remote database
func (db *firestoreDB) disconnect() {
	// Close the client connection if it is open
	if db.client != nil {
		defer db.client.Close()
	}
}

// Connect allows the user to connect to the database
func (db *firestoreDB) connect() {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("FIRESTORE_PROJECT_ID")

	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		panic(err)
	}

	// Etablish Database Collection object
	db.client = client
	db.games = db.client.Collection("games")
	db.players = db.client.Collection("players")
}

func init() {
	registerDB(&DB{
		name:        "FIRESTORE",
		description: "Production Firestore connection",
		UnoDB:       new(firestoreDB),
	})
}
