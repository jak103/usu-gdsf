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

type mongoDB struct {
	client      *mongo.Client
	database    *mongo.Database
	gameRecords *mongo.Collection
}

func (db *mongoDB) GetAllGameRecords() (*[]models.GameRecord, error) {
	games := make([]models.GameRecord, 0)

	cursor, err := db.gameRecords.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		g := models.GameRecord{}
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
	if err = client.Connect(ctx); err != nil {
		log.Warn("Unable to establish database connection.")
		return
	}
	db.client = client
	database := client.Database("usu-gdsf")
	db.database = database
	db.gameRecords = database.Collection("gameRecords")

	if count, err := db.gameRecords.CountDocuments(ctx, bson.D{{}}); err != nil {
		log.Error("There was a problem getting the documents from the Games Record collection: %v", err)
	} else if count == 0 {
		log.Debug("No game records currently exist. Seeding the games record collection...")

		docs := []interface{}{}

		for _, v := range CreateGamesFromJson() {
			doc, err := bson.Marshal(v)
			if err != nil {
				log.Error("Error occurred while creating document: %v", err)
				return
			}
			docs = append(docs, doc)
		}

		if insertManyResult, insertErr := db.gameRecords.InsertMany(ctx, docs); insertErr != nil {
			log.Error("An error happened while seeding the collection: %v", insertErr)
		} else {
			log.Debug("Inserted multiple documents: ", insertManyResult.InsertedIDs)
		}
	}
}

func init() {
	registerDB(&DB{
		Name:          "MONGO",
		Description:   "Mongo database for dev connections",
		StoreDatabase: new(mongoDB),
	})
}
