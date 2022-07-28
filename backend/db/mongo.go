package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jak103/usu-gdsf/config"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var _ Database = (*Mongo)(nil)

type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
	games    *mongo.Collection
}

func (db Mongo) GetGameID(game models.Game) (string, error) {
	result := db.database.Collection("games").FindOne(context.Background(), bson.M{
		// TODO might need to convert game datatypes to mongo datatypes
		"name":         game.Name,
		"author":       game.Author,
		"creationDate": game.CreationDate,
		"version":      game.Version,
	}, options.FindOne().SetShowRecordID(true))

	// find error
	if result.Err() == mongo.ErrNoDocuments {
		log.Error("No document found in Mongo GetGameID")
		return "", result.Err()
	}

	// decode found document
	raw, err := result.DecodeBytes()
	if err != nil {
		log.WithError(err).Error("Cannot decode result in Mongo GetGameID")
		return "", err
	}

	// lookup the id in the raw bson data
	id := raw.Lookup("recordId").String()
	if id == "" {
		log.Error("Could not find ID from db data in Mongo GetGameID")
		return id, errors.New("could not find 'recordId' key in raw Mongo data")
	}
	// TODO might need to convert id to hex id
	return id, nil
}

func (db Mongo) GetGameByID(id string) (models.Game, error) {
	// convert hex id to object ID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.WithError(err).Error("Invalid id in Mongo ID search")
	}

	// find game with object ID
	result := db.database.Collection("games").FindOne(context.Background(), bson.M{"_id": objID})
	game := models.Game{}
	decodeErr := result.Decode(game)
	if decodeErr != nil {
		log.WithError(decodeErr).Error("Mongo decoding Game struct error")
		return game, decodeErr
	}

	return game, nil
}

func (db Mongo) AddGame(game models.Game) (string, error) {
	insertResult, err := db.database.Collection("games").InsertOne(context.Background(), game)
	if err != nil {
		log.WithError(err).Error("Failed to add game to Mongo db")
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
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
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUri))
	if err != nil {
		log.WithError(err).Error("Failed to create mongo client")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		log.WithError(err).Warn("Unable to establish database connection.")
		return err
	}
	db.client = client
	database := client.Database("usu-gdsf")
	db.database = database
	db.games = database.Collection("games")

	// Logic for creating seed data
	if count, err := db.games.CountDocuments(ctx, bson.D{{}}); err != nil {
		log.Error("There was a problem getting the documents from the Games Record collection: %v", err)
	} else if count == 0 {
		log.Debug("No game records currently exist. Seeding the games record collection...")

		docs := []interface{}{}

		for _, v := range CreateGamesFromJson() {
			doc, err := bson.Marshal(v)
			if err != nil {
				log.Error("Error occurred while creating document: %v", err)
				return err
			}
			docs = append(docs, doc)
		}

		if insertManyResult, insertErr := db.games.InsertMany(ctx, docs); insertErr != nil {
			log.Error("An error happened while seeding the collection: %v", insertErr)
		} else {
			log.Debug("Inserted multiple documents: ", insertManyResult.InsertedIDs)
		}
	}

	return nil
}
