package db

import (
	"context"
	"errors"
	"fmt"
	"math"
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
	primitiveObjectId, err := primitive.ObjectIDFromHex(game.Id)
	if err != nil {
		log.WithError(err).Error("error on getting primitive object id from hex string")
		return err
	}

	gc := db.database.Collection("games")
	res, err := gc.DeleteOne(context.Background(), bson.M{
		"_id": primitiveObjectId,
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

func (db Mongo) RemoveGameByTag(tag string) error {
	gc := db.database.Collection("games")
	_, err := gc.DeleteMany(context.Background(), bson.M{"tags": tag})
	if err != nil {
		log.WithError(err).Error("error on removing collections with given tag")
		return err
	}
	return nil
}

func (db Mongo) SortGames(field_name string, order int) ([]models.Game, error) {
	order1 := float64(order)
	if math.Abs(order1) != 1 {
		log.Error("sorting order is not correct")
		err1 := errors.New("sorting order is not among -1 and 1")
		return nil, err1
	}
	gc := db.database.Collection("games")
	options := options.Find()
	options.SetSort(bson.D{{field_name, order}})
	options.SetLimit(10)
	cursor, err := gc.Find(context.Background(), bson.D{}, options)
	// it does not need to close the cursor in this case but just for sanity
	if err != nil {
		log.WithError(err).Error("couldn't complete the sorting query")
	}
	defer cursor.Close(context.Background())

	results := make([]models.Game, 0)
	for cursor.Next(context.Background()) {
		// create a value into which the single document can be decoded
		var gameObject models.Game
		err := cursor.Decode(&gameObject)
		if err != nil {
			log.WithError(err).Error("couldn't decode the cursor")
			return nil, err
		}
		results = append(results, gameObject)

	}

	if err := cursor.Err(); err != nil {
		log.WithError(err).Error("last error on reading cursor")
		return nil, err
	}
	return results, nil
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

// GetGamesByTags search and return all games with given tag
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

// GroupGamesByFirstLetter and return all games starting with the given letter

func (db Mongo) GetGamesByFirstLetter(letter string) ([]models.Game, error) {
	if len(letter) != 1 {
		log.Error("Please enter a valid character")
		err1 := errors.New("Not a valid charater entered")
		return nil, err1
	}
	gc := db.database.Collection("games")
	cur, err := gc.Find(context.Background(), bson.M{"name": bson.M{"$regex": "^" + letter, "$options": "i"}})
	if err != nil {
		log.WithError(err).Error("Error getting games with Firest Letter")
		return nil, err
	}
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
	game, _ := DecodeGameBsonData(data)
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

// Add download history object and to return its ID
func (db Mongo) AddDownload(download models.Download) (string, error) {
	// Logic to check if one already exists

	insertResult, err := db.database.Collection("downloads").InsertOne(context.Background(), download)
	if err != nil {
		log.WithError(err).Error("Failed to add download to Mongo db")
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (db Mongo) GetDownloadByID(id string) (models.Download, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.WithError(err).Error("Invalid id in Mongo ID search")
	}

	result := db.database.Collection("downloads").FindOne(context.Background(), bson.M{"_id": objID})
	data := bson.M{}

	err = result.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Cannot decode record in Mongo GetDownloadByID")
		return models.Download{}, err
	}

	download, _ := DecodeDownloadBsonData(data)
	return download, nil
}

func DecodeCursorToGame(cur *mongo.Cursor) (models.Game, error) {
	data := bson.M{}
	err := cur.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Unable to decode Mongo game cursor")
		return models.Game{}, err
	}

	return DecodeGameBsonData(data)
}

func DecodeCursorToDownload(cur *mongo.Cursor) (models.Download, error) {
	data := bson.M{}
	err := cur.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Unable to decode Mongo download cursor")
		return models.Download{}, err
	}

	return DecodeDownloadBsonData(data)
}

func convert[T any](v any) T {
	if v == nil {
		return *new(T)
	}
	return v.(T)
}

// Method used to decode data that is shared across object times. Including tags, and creationDate
func DecodeCommonData(data bson.M) ([]string, time.Time, error) {
	var err error

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

	// Check the type of creationDate date (time.Time or string)
	if _, ok := data["creationdate"].(primitive.DateTime); ok { // creationDate is saved as time.Time
		date = data["creationdate"].(primitive.DateTime).Time().UTC()
	} else if _, ok := data["creationdate"].(string); ok { // creationDate is saved as string
		date, err = time.Parse("1/2/2006", data["creationdate"].(string))
		if err != nil {
			log.WithError(err).Error("Cannot parse string into date in Mongo GetGameByID")
			var emptyTags []string
			var blankDate time.Time
			return emptyTags, blankDate, err
		}
	}

	return tags, date, nil
}

func DecodeDownloadBsonData(data bson.M) (models.Download, error) {
	_, date, err := DecodeCommonData(data)

	download := models.Download{
		Id:           data["_id"].(primitive.ObjectID).Hex(),
		UserId:       convert[string](data["userid"]),
		GameId:       convert[string](data["gameid"]),
		CreationDate: date,
	}

	if err != nil {
		log.WithError(err).Error("Cannot Decode Download Object")

	}

	return download, nil
}

func DecodeGameBsonData(data bson.M) (models.Game, error) {
	tags, date, err := DecodeCommonData(data)

	// load game model
	game := models.Game{
		Id:           data["_id"].(primitive.ObjectID).Hex(),
		Name:         convert[string](data["name"]),
		Rating:       float32(convert[float64](data["rating"])),
		TimesPlayed:  int(convert[int32](data["timesplayed"])),
		ImagePath:    convert[string](data["imagepath"]),
		Description:  convert[string](data["description"]),
		Developer:    convert[string](data["developer"]),
		CreationDate: date,
		Version:      convert[string](data["version"]),
		Tags:         tags,
		Downloads:    convert[int64](data["downloads"]),
		DownloadLink: convert[string](data["downloadlink"]),
	}

	if err != nil {
		log.WithError(err).Error("Cannot Decode Game Object")
	}

	return game, nil
}

func (db Mongo) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)

	gc := db.database.Collection("games")
	cursor, err := gc.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		log.WithError(err).Error("mongo game find failed")
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

func (db Mongo) GetAllDownloads() ([]models.Download, error) {
	downloads := make([]models.Download, 0)

	gc := db.database.Collection("downloads")
	cursor, err := gc.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		log.WithError(err).Error("mongo download find failed")
		return nil, err
	}

	for cursor.Next(context.Background()) {
		d, err := DecodeCursorToDownload(cursor)
		if err != nil {
			return nil, err
		}
		downloads = append(downloads, d)
	}

	return downloads, nil
}

func (db Mongo) CreateUser(newUser models.User) (models.User, error) {
	// users := db.database.Collection("users")

	// newUserDoc, err := users.InsertOne(context.Background(), newUser, nil)
	// if err != nil {
	// 	log.WithError(err).Error("Failed to insert new user")
	// 	return nil, err
	// }

	return newUser, nil
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
	// if database and collection does not exist it will create one
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

// add a new review and recalculate its average
func (db Mongo) AddReview(review models.Review) (string, error) {
	insertResult, err := db.database.Collection("reviews").InsertOne(context.Background(), review)
	if err != nil {
		log.WithError(err).Error("Failed to add review to Mongo db")
	}

	// recalculate the average review score
	game, err := db.GetGameByID(review.GameId)
	game.ReviewIds = append(game.ReviewIds, review.Id)

	totalScore := game.AverageReview * (float64((len(game.ReviewIds) - 1)))
	game.AverageReview = ((totalScore + review.Score) / float64(len(game.ReviewIds)))
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (db Mongo) GetReviewByID(id string) (models.Review, error) {
	// convert hex id to object ID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.WithError(err).Error("Invalid id in Mongo ID search")
	}

	// find review with object ID
	result := db.database.Collection("reviews").FindOne(context.Background(), bson.M{"_id": objID})

	// decode into bson
	data := bson.M{}
	err = result.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Cannot decode record in Mongo GetReviewByID")
		return models.Review{}, err
	}

	// decode bson into review
	review, _ := DecodeBsonReviewData(data)
	return review, nil
}

func DecodeBsonReviewData(data bson.M) (models.Review, error) {
	// decode tags array
	var tags []string
	if data["tags"] != nil {
		primTags := data["tags"].(primitive.A)
		tags = make([]string, len(primTags))
		for i, v := range primTags {
			tags[i] = v.(string)
		}
	}

	// load review model
	review := models.Review{
		Id:     data["_id"].(primitive.ObjectID).Hex(),
		GameId: data["GameId"].(string),
		UserId: data["UserId"].(uint64),
		Score:  data["Score"].(float64),
		Text:   data["Text"].(string),
	}

	return review, nil
}

// RemoveReview removes the given Review from the db
func (db Mongo) RemoveReview(review models.Review) error {
	primitiveObjectId, err := primitive.ObjectIDFromHex(review.Id)
	if err != nil {
		log.WithError(err).Error("error on getting primitive object id from hex string")
		return err
	}

	gc := db.database.Collection("reviews")
	res, err := gc.DeleteOne(context.Background(), bson.M{
		"_id": primitiveObjectId,
	})

	if err != nil {
		log.WithError(err).Error("Mongo RemoveReview deletion error")
		return err
	}

	if res.DeletedCount > 1 {
		log.Error("Mongo RemoveReview deleted more than one record")
		return errors.New("mongo deleted more than one record")
	}

	return nil
}
