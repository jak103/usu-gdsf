package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
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
	games    *mongo.Collection
}

func (db Mongo) GetGameByID(id uuid.UUID) (*models.Game, error) {
	uuidID, err := uuid.Parse(id.String())
	if err != nil {
		return nil, err
	}
	var res models.Game
	if err := db.games.FindOne(context.Background(), bson.M{"id": uuidID}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (db Mongo) GetGamesByTags(tags []string) ([]models.Game, error) {
	games := make([]models.Game, 0)
	cursor, err := db.games.Find(context.Background(), bson.M{"tags": bson.M{"$in": tags}})
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

func (d *Mongo) GetGamesByPublishDate(startRange string, endRange string) ([]models.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) CreateGame(newGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) DeleteGame(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) UpdateGame(updatedGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

// Users
func (d *Mongo) GetAllUsers() ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}


func (d *Mongo) GetUserByID(id uuid.UUID)(*models.User, error) {

	panic("not implemented") // TODO: Implement
}

func (d *Mongo) GetUsersByRole(role int64) ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) CreateUser(newUser models.User) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) DeleteUser(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) UpdateUser(updatedUser models.User) error {
	panic("not implemented") // TODO: Implement
}

// Ratings
func (d *Mongo) GetRatingByID(id uuid.UUID) (*models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) GetRatingsByGame(gameID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) GetRatingsByUser(userID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) CreateRating(newRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) DeleteRating(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) DeleteRatingsByGame(gameID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mongo) updateRating(updatedRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
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
