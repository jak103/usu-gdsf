package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jak103/usu-gdsf/config"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ Database = (*Mongo)(nil)

type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
	games    *mongo.Collection
}

// RemoveGame removes the given game from the db
func (db Mongo) RemoveGame(game models.Game) error {
	gc := db.database.Collection("games")
	res, err := gc.DeleteOne(context.Background(), bson.M{
		"name":         game.Name,
		"rating":       game.Rating,
		"timesplayed":  game.TimesPlayed,
		"imagepath":    game.ImagePath,
		"description":  game.Description,
		"developer":    game.Developer,
		"creationdate": game.CreationDate,
		"version":      game.Version,
		"tags":         game.Tags,
		"downloads":    game.Downloads,
		"downloadlink": game.DownloadLink,
	})
	if err != nil {
		log.WithError(err).Error("Mongo RemoveGame deletion error")
		return err
	}

	if res.DeletedCount > 1 {
		log.Error("Mongo RemoveGame deleted more than one record")
		return errors.New("mongo deleted more than one record")
	}

	return nil
}

// GetGamesByTag search and return all games with given tag
func (db Mongo) GetGamesByTag(s string) ([]models.Game, error) {
	// Search for games containing tag
	gc := db.database.Collection("games")
	cur, err := gc.Find(context.Background(), bson.D{
		{"tags", s},
	})
	if err != nil {
		log.WithError(err).Error("Mongo GetGamesByTag search error")
		return nil, err
	}

	// decode games containing tag
	games := make([]models.Game, 0)
	for cur.Next(context.Background()) {
		g, err := DecodeCursorToGame(cur)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}

// GetGamesByTag search and return all games with given tag
func (db Mongo) GetGamesByTags(tags []string, matchAll bool) ([]models.Game, error) {
	result, err := db.GetGamesByTag(tags[0])
	if err != nil {
		log.WithError(err).Error("Error getting games with tags")
		return nil, err
	}

	for _, tag := range tags[1:] {
		games, err := db.GetGamesByTag(tag)

		if err != nil {
			log.WithError(err).Error("Error getting games with tags")
			return nil, err
		}

		if matchAll {
			for i, game := range result {
				if !containsGame(games, game) {
					result[i] = result[len(result)-1]
					result = result[:len(result)-1]
				}
			}
		} else {
			for _, game := range games {
				if !containsGame(result, game) {
					result = append(result, game)
				}
			}
		}
	}
	return result, nil
}

// Helper function to check if one array contains an element
func containsGame(games []models.Game, game models.Game) bool {
	for _, v := range games {
		if v.Id == game.Id {
			return true
		}
	}
	return false
}

// GetGameByID find and return the game with the given db hex id
func (db Mongo) GetGameByID(id string) (models.Game, error) {
	// convert hex id to object ID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.WithError(err).Error("Invalid id in Mongo ID search")
	}

	// find game with object ID
	result := db.database.Collection("games").FindOne(context.Background(), bson.M{"_id": objID})

	// decode into bson
	data := bson.M{}
	err = result.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Cannot decode record in Mongo GetGameByID")
		return models.Game{}, err
	}

	// decode bson into game
	game, _ := DecodeBsonData(data)
	return game, nil
}

// AddGame add game to database. Returns assigned ID
func (db Mongo) AddGame(game models.Game) (string, error) {
	insertResult, err := db.database.Collection("games").InsertOne(context.Background(), game)
	if err != nil {
		log.WithError(err).Error("Failed to add game to Mongo db")
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func DecodeCursorToGame(cur *mongo.Cursor) (models.Game, error) {
	data := bson.M{}
	err := cur.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Unable to decode Mongo cursor")
		return models.Game{}, err
	}

	return DecodeBsonData(data)
}

func DecodeBsonData(data bson.M) (models.Game, error) {
	// decode tags array
	var tags []string
	if data["tags"] != nil {
		primTags := data["tags"].(primitive.A)
		tags = make([]string, len(primTags))
		for i, v := range primTags {
			tags[i] = v.(string)
		}
	}

	//decode creationDate
	var date time.Time
	var err error
	// Check the type of creationDate date (time.Time or string)
	if _, ok := data["creationdate"].(primitive.DateTime); ok { // creationDate is saved as time.Time
		date = data["creationdate"].(primitive.DateTime).Time().UTC()
	} else if _, ok := data["creationdate"].(string); ok { // creationDate is saved as string
		date, err = time.Parse("1/2/2006", data["creationdate"].(string))
		if err != nil {
			log.WithError(err).Error("Cannot parse string into date in Mongo GetGameByID")
			return models.Game{}, err
		}
	}

	// load game model
	game := models.Game{
		Id:           data["_id"].(primitive.ObjectID).Hex(),
		Name:         data["name"].(string),
		Rating:       float32(data["rating"].(float64)),
		TimesPlayed:  int(data["timesplayed"].(int32)),
		ImagePath:    data["imagepath"].(string),
		Description:  data["description"].(string),
		Developer:    data["developer"].(string),
		CreationDate: date,
		Version:      data["version"].(string),
		Tags:         tags,
		Downloads:    data["downloads"].(int64),
		DownloadLink: data["downloadlink"].(string),
	}

	return game, nil
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
		g, err := DecodeCursorToGame(cursor)
		if err != nil {
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
