package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ Database = (*Mongo)(nil)

type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
}

func (db Mongo) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)

	gc := db.database.Collection("games")
	cursor, err := gc.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		log.WithError(err).Error("mongo find failed")
		return nil, err
	}

	for cursor.Next(context.Background()) {
		g := models.Game{}
		err := cursor.Decode(&g)
		if err != nil {
			log.WithError(err).Error("Failed to decode cursor")
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}

// disconnect disconnects from the remote database
func (db *Mongo) Disconnect() error {
	fmt.Println("Disconnecting from the database.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.client.Disconnect(ctx); err != nil {
		log.WithError(err).Error("Failed to disconnect from mongo")
		return err
	}

	return nil
}

// connect allows the user to connect to the database
func (db *Mongo) Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.WithError(err).Error("Failed to create mongo client")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to connect to mongo")
		return err
	}

	db.client = client
	database := client.Database("gdsf") // TODO This should probably be a parameter
	db.database = database

	return nil
}
