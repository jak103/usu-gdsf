package db

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

var _ Database = (*Mock)(nil)

type Mock struct {
	games   map[uuid.UUID]models.Game
	users   map[uuid.UUID]models.User
	ratings map[uuid.UUID]models.GameRating
}

func (db *Mock) GetGameByID(id uuid.UUID) (*models.Game, error) {
	if id.String() != "" {
		if game, ok := db.games[id]; ok {
			return &game, nil
		}
	}
	return nil, errors.New("mockdb: game not found")
}

func (db *Mock) GetGamesByTags(tags []string) ([]models.Game, error) {
	games := make(map[string]models.Game)

	for _, game := range db.games {
		for _, tag := range tags {
			valid := map[string]bool{}
			game_tags := make([]string, 0)
			for _, v := range game_tags {
				valid[v] = true
			}
			if valid[tag] {
				games[game.ID.String()] = game
			}
		}
	}
	game_slice := []models.Game{}
	for _, game := range games {
		game_slice = append(game_slice, game)
	}

	return game_slice, nil
}

func (db *Mock) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)

	for _, game := range db.games {
		games = append(games, game)
	}

	return games, nil
}

func (d *Mock) GetGamesByPublishDate(startRange string, endRange string) ([]models.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) CreateGame(newGame models.Game) error {
	if newGame.ID == uuid.Nil {
		log.Error("newGame struct has nil ID")
		return errors.New("newGame struct has nil ID")
	}
	if _, exists := d.games[newGame.ID]; !exists {
		d.games[newGame.ID] = newGame
		return nil
	}
	log.Error("newGame ID already exists in mock db")
	return errors.New("game already exists in mock db")
}

func (d *Mock) DeleteGame(id uuid.UUID) error {
	if id == uuid.Nil {
		log.Error("nil id provided")
		return errors.New("nil id provided")
	}
	if _, exists := d.games[id]; exists {
		delete(d.games, id)
		return nil
	}
	log.Error("provided id doesn't exist in db")
	return errors.New("provided id doesn't exist in db")
}

func (d *Mock) UpdateGame(updatedGame models.Game) error {
	if updatedGame.ID == uuid.Nil {
		log.Error("updatedGame struct has nil ID")
		return errors.New("updatedGame struct has nil ID")
	}
	if _, exists := d.games[updatedGame.ID]; !exists {
		log.Error("updatedGame ID does not exists in mock db")
		return errors.New("updatedGame ID does not exists in mock db")

	}
	d.games[updatedGame.ID] = updatedGame
	return nil
}

// Users
func (d *Mock) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)

	for _, user := range d.users {
		users = append(users, user)
	}

	return users, nil
}

func (db *Mock) GetUserByID(id uuid.UUID) (*models.User, error) {
	if id == uuid.Nil {
		log.Error("nil id provided")
		return nil, errors.New("nil id provided")
	}
	if user, ok := db.users[id]; ok {
		return &user, nil
	}
	return nil, errors.New("mockdb: user not found")
}

func (d *Mock) GetUsersByRole(role int64) ([]models.User, error) {
	if role != 0 && role != 1 {
		return nil, errors.New(fmt.Sprintf("User role %v does not exist", role))
	}
	users := make([]models.User, 0)
	queryRole := models.Role(role)
	for _, user := range d.users {
		if user.Role == queryRole {
			users = append(users, user)
		}
	}
	return users, nil
}

func (d *Mock) CreateUser(newUser models.User) error {
	if newUser.ID == uuid.Nil {
		log.Error("newUser struct has nil ID")
		return errors.New("newUser struct has nil ID")
	}
	if _, exists := d.users[newUser.ID]; !exists {
		d.users[newUser.ID] = newUser
		return nil
	}
	log.Error("newUser ID already exists in mock db")
	return errors.New("newUser already exists in mock db")
}

func (d *Mock) DeleteUser(id uuid.UUID) error {
	if id == uuid.Nil {
		log.Error("nil id provided")
		return errors.New("nil id provided")
	}
	if _, exists := d.users[id]; exists {
		delete(d.users, id)
		return nil
	}
	log.Error("provided user id doesn't exist in db")
	return errors.New("provided user id doesn't exist in db")
}

func (d *Mock) UpdateUser(updatedUser models.User) error {
	if updatedUser.ID == uuid.Nil {
		log.Error("updatedUser struct has nil ID")
		return errors.New("updatedUser struct has nil ID")
	}
	if _, exists := d.users[updatedUser.ID]; !exists {
		log.Error("updatedUser ID does not exists in mock db")
		return errors.New("updatedUser ID does not exists in mock db")

	}
	d.users[updatedUser.ID] = updatedUser
	return nil
}

// Ratings
func (d *Mock) GetRatingByID(id uuid.UUID) (*models.GameRating, error) {
	if id == uuid.Nil {
		log.Error("nil rating id provided")
		return nil, errors.New("nil rating id provided")
	}
	if rating, ok := d.ratings[id]; ok {
		return &rating, nil
	}
	return nil, errors.New("mockdb: rating not found")
}

func (d *Mock) GetRatingsByGame(gameID uuid.UUID) ([]models.GameRating, error) {
	if gameID == uuid.Nil {
		log.Error("nil game id provided")
		return nil, errors.New("nil game id provided")
	}
	ratings := make([]models.GameRating, 0)

	for _, rating := range d.ratings {
		if rating.GameId == gameID {
			ratings = append(ratings, rating)
		}
	}

	return ratings, nil
}

func (d *Mock) GetRatingsByUser(userID uuid.UUID) ([]models.GameRating, error) {
	if userID == uuid.Nil {
		log.Error("nil user id provided")
		return nil, errors.New("nil user id provided")
	}
	ratings := make([]models.GameRating, 0)

	for _, rating := range d.ratings {
		if rating.UserID == userID {
			ratings = append(ratings, rating)
		}
	}

	return ratings, nil
}

func (d *Mock) CreateRating(newRating models.GameRating) error {
	if newRating.ID == uuid.Nil {
		log.Error("newRating struct has nil ID")
		return errors.New("newRating struct has nil ID")
	}
	if _, exists := d.ratings[newRating.ID]; !exists {
		d.ratings[newRating.ID] = newRating
		return nil
	}
	log.Error("newRating ID already exists in mock db")
	return errors.New("newRating already exists in mock db")
}

func (d *Mock) DeleteRating(id uuid.UUID) error {
	if id == uuid.Nil {
		log.Error("nil rating id provided")
		return errors.New("nil rating id provided")
	}
	if _, exists := d.ratings[id]; exists {
		delete(d.ratings, id)
		return nil
	}
	log.Error("provided rating id doesn't exist in db")
	return errors.New("provided rating id doesn't exist in db")
}

func (d *Mock) DeleteRatingsByGame(gameID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) UpdateRating(updatedRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
}

func (db *Mock) Connect() error {
	log.Info("mock connect")

	if len(db.games) == 0 {
		games := make(map[uuid.UUID]models.Game)
		for _, v := range CreateGamesFromJson() {
			games[v.ID] = v
		}
		db.games = games
	}

	if len(db.users) == 0 {
		users := make(map[uuid.UUID]models.User)
		db.users = users
	}

	log.Debug("Sample Database Initialized")

	return nil
}

func (db Mock) Disconnect() error {
	return nil
}

// func init(){

// }
