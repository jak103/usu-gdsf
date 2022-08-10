package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"go.uber.org/multierr"
	"google.golang.org/api/iterator"
)

var _ Database = (*Firestore)(nil)

type Firestore struct {
	client  *firestore.Client
	games   *firestore.CollectionRef
	users   *firestore.CollectionRef
	ratings *firestore.CollectionRef
}

func (db Firestore) GetGameByID(id uuid.UUID) (*models.Game, error) {
	gameDoc := db.games.Doc(id.String())
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

func (db Firestore) GetGamesByTags(tags []string) ([]models.Game, error) {
	games := make([]models.Game, 0)
	gc := db.client.Collection("games")

	documents := gc.Where("tags", "in", tags).Documents(context.Background())

	for {
		docRef, docRefErr := documents.Next()

		if docRefErr == iterator.Done {
			break
		}

		var game models.Game
		docRef.DataTo(&game)

		games = append(games, game)
	}

	return games, nil
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

// func (d *Firestore) GetGamesByTags(tags []string) ([]models.Game, error) {
// 	panic("not implemented") // TODO: Implement
// }

func (db *Firestore) GetGamesByPublishDate(startRange string, endRange string) ([]models.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) CreateGame(newGame models.Game) error {
	_, err := db.client.Collection("games").Doc(newGame.ID.String()).Set(context.Background(), newGame)
	if err != nil {
		log.Error("firestore: could not create game")
		return err
	}
	return nil
}

func (db *Firestore) DeleteGame(id uuid.UUID) error {
	snapshot, err := db.client.Collection("games").Doc(id.String()).Get(context.Background())
	if err != nil {
		log.Error("firestore: could not get game snapshot")
		return err
	}

	_, err = snapshot.Ref.Delete(context.Background())
	if err != nil {
		log.Error("firestore: could not delete game snapshot")
		return err
	}
	return nil
}

func (db *Firestore) UpdateGame(updatedGame models.Game) error {
	snapshot, err := db.client.Collection("games").Doc(updatedGame.ID.String()).Get(context.Background())
	if err != nil {
		log.Error("firestore: could not get game snapshot")
		return err
	}
	gameValues := reflect.ValueOf(updatedGame)
	gameType := gameValues.Type()
	for i := 0; i < gameValues.NumField(); i++ {
		_, err = snapshot.Ref.Update(context.Background(), []firestore.Update{{
			Path:  gameType.Field(i).Name,
			Value: gameValues.Field(i).Interface(),
		}})
		if err != nil {
			log.Error(fmt.Sprintf("firestore: could not update game field: %v", gameType.Field(i).Name))
			return err
		}
	}
	return nil
}

// Users
func (db *Firestore) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	uc := db.client.Collection("users")

	documents := uc.DocumentRefs(context.Background())
	for {
		docRef, docRefErr := documents.Next()

		if docRefErr == iterator.Done {
			break
		}

		var user models.User

		if docSnapshot, _ := docRef.Get(context.Background()); docSnapshot != nil {
			_ = docSnapshot.DataTo(&user)
		}

		users = append(users, user)
	}

	return users, nil
}

func (db *Firestore) GetUserByID(id uuid.UUID) (*models.User, error) {
	userDoc := db.users.Doc(id.String())
	docSnapshot, err := userDoc.Get(context.Background())

	if err != nil {
		return nil, err
	}

	if docSnapshot == nil {
		return nil, fmt.Errorf("%s: user not found", id)
	}

	var user models.User
	if err = docSnapshot.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Username is not currently unique
func (db *Firestore) GetUserByUserName(userName string) (*models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) GetUsersByRole(role int64) ([]models.User, error) {
	users := make([]models.User, 0)
	gc := db.client.Collection("users")

	documents := gc.Where("role", "==", role).Documents(context.Background())

	for {
		docRef, docRefErr := documents.Next()

		if docRefErr == iterator.Done {
			break
		}

		var user models.User
		docRef.DataTo(&user)

		users = append(users, user)
	}

	return users, nil
}

func (db *Firestore) CreateUser(newUser models.User) error {
	if newUser.ID == uuid.Nil {
		log.Error("newUser struct has nil ID")
		return errors.New("newUser struct has nil ID")
	}

	uc := db.client.Collection("users")
	exists := uc.Where("id", "==", newUser.ID).Documents(context.Background())

	if exists == nil {
		uc.Doc(newUser.ID.String()).Set(context.Background(), newUser)
		return nil
	}

	fmt.Println(exists.GetAll())

	log.Error("newUser ID already exists in firestore_emu db")
	return errors.New("newUser already exists in firestore_emu db")
}

func (db *Firestore) DeleteUser(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (db *Firestore) UpdateUser(updatedUser models.User) error {
	panic("not implemented") // TODO: Implement
}

// Ratings
func (db *Firestore) GetRatingByID(id uuid.UUID) (*models.GameRating, error) {
	ratingDoc := db.ratings.Doc(id.String())
	snapshot, err := ratingDoc.Get(context.Background())

	if err != nil {
		return nil, err
	}

	if snapshot == nil {
		return nil, fmt.Errorf("%s: rating not found", id)
	}

	var rating models.GameRating
	if err = snapshot.DataTo(&rating); err != nil {
		return nil, err
	}

	return &rating, nil
}

func (db *Firestore) GetRatingsByGame(gameID uuid.UUID) ([]models.GameRating, error) {
	ratings := make([]models.GameRating, 0)
	ratingsCollection := db.ratings

	docs := ratingsCollection.Where("gameid", "==", gameID.String()).Documents(context.Background())

	for {
		docRef, docRefErr := docs.Next()

		if docRefErr == iterator.Done {
			break
		}

		var rating models.GameRating
		docRef.DataTo(&rating)

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

func (db *Firestore) GetRatingsByUser(userID uuid.UUID) ([]models.GameRating, error) {
	ratings := make([]models.GameRating, 0)
	ratingsCollection := db.ratings

	docs := ratingsCollection.Where("userid", "==", userID.String()).Documents(context.Background())

	for {
		docRef, docRefErr := docs.Next()

		if docRefErr == iterator.Done {
			break
		}

		var rating models.GameRating
		docRef.DataTo(&rating)

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

func (db *Firestore) CreateRating(newRating models.GameRating) error {
	// NOTE: this assumes that the ID is already set in the new object
	_, err := db.ratings.Doc(newRating.ID.String()).Set(context.Background(), newRating)
	if err != nil {
		log.WithError(err).Error("failed to add a rating to firestore")
	}
	return err
}

func (db *Firestore) DeleteRating(id uuid.UUID) error {
	ratingDoc := db.ratings.Doc(id.String())
	_, err := ratingDoc.Delete(context.Background())

	return err
}

func (db *Firestore) DeleteRatingsByGame(gameID uuid.UUID) error {
	ratingsCollection := db.ratings

	docs := ratingsCollection.Where("gameid", "==", gameID.String()).Documents(context.Background())

	var finalErr error = nil
	for {
		docRef, docRefErr := docs.Next()
		if docRefErr == iterator.Done {
			break
		}
		_, err := docRef.Ref.Delete(context.Background())
		if err != nil {
			if finalErr == nil {
				finalErr = err
			} else {
				finalErr = multierr.Append(finalErr, err)
			}
		}
	}

	return finalErr
}

func (db *Firestore) UpdateRating(updatedRating models.GameRating) error {
	var finalErr error = nil
	_, getErr := db.ratings.Doc(updatedRating.ID.String()).Get(context.Background())
	if getErr == nil {
		_, err := db.ratings.Doc(updatedRating.ID.String()).Set(context.Background(), updatedRating)
		if err != nil {
			log.WithError(err).Error("failed to update a rating in firestore")
			finalErr = err
		}
	} else {
		log.WithError(getErr).Error("Tried to update a rating that doesn't exist")
		finalErr = getErr
	}

	return finalErr
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
	db.games = client.Collection(GAMES)
	db.users = client.Collection(USERS)
	db.ratings = client.Collection(RATINGS)

	return nil
}
