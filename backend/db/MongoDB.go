package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jak103/uno/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client   *mongo.Client
	uri      string
	database *mongo.Database
	games    *mongo.Collection
	players  *mongo.Collection
}

func (db *mongoDB) GetAllGames() (*[]model.Game, error) {
	games := make([]model.Game, 0)

	cursor, err := db.games.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		g := model.Game{}
		err := cursor.Decode(&g)
		if err != nil {
			panic(err)
		}
		games = append(games, g)
	}

	return &games, nil
}

// disconnect disconnects from the remote database
func (db *mongoDB) disconnect() {
	fmt.Println("Disconnecting from the database.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.client.Disconnect(ctx); err != nil {
		panic(err)
	}
	defer cancel()
}

// connect allows the user to connect to the database
func (db *mongoDB) connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db.client = client
	database := client.Database("uno")
	db.database = database
	db.games = database.Collection("games")
	db.players = database.Collection("players")
}

func init() {
	registerDB(&DB{
		name:        "MONGO",
		description: "Mongo database for dev connections",
		GdsfDB:      new(mongoDB),
	})
}
