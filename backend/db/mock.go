package db

import (
	"errors"

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

func (db *Mock) GetUserByUserName(userName string) (*models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Mock) GetUsersByRole(role int64) ([]models.User, error) {
	typeRole := models.Role(role)
	user_slice := []models.User{}
	for _, user := range db.users {
		if user.Role == typeRole {
			user_slice = append(user_slice, user)
		}

	}

	return user_slice, nil
}

func (db *Mock) CreateUser(newUser models.User) error {
	if newUser.ID == uuid.Nil {
		log.Error("newUser struct has nil ID")
		return errors.New("newUser struct has nil ID")
	}

	if _, exists := db.users[newUser.ID]; !exists {
		db.users[newUser.ID] = newUser
		return nil
	}
	log.Error("newUser ID already exists in mock db")
	return errors.New("newUser already exists in mock db")
}

func (db *Mock) DeleteUser(id uuid.UUID) error {
	if id == uuid.Nil {
		log.Error("nil id provided")
		return errors.New("nil id provided")
	}

	if _, exists := db.users[id]; exists {
		delete(db.users, id)
		return nil
	}
	log.Error("provided id doesn't exist in db")
	return errors.New("provided id doesn't exist in db")
}

func (db *Mock) UpdateUser(updatedUser models.User) error {
	if updatedUser.ID == uuid.Nil {
		log.Error("updatedUser struct has nil ID")
		return errors.New("updatedUser struct has nil ID")
	}

	if _, exists := db.games[updatedUser.ID]; !exists {
		log.Error("updatedUser ID does not exist in mock db")
		return errors.New("updatedUser ID does not exist in mock db")
	}
	db.users[updatedUser.ID] = updatedUser
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
	ratings, err := d.GetRatingsByGame(gameID)
	if ratings == nil {
		return err
	}
	for _, rating := range ratings {
		delete(d.ratings, rating.ID)
	}
	return nil
}

func (d *Mock) UpdateRating(updatedRating models.GameRating) error {
	if updatedRating.ID == uuid.Nil {
		log.Error("updatedRating struct has nil ID")
		return errors.New("updatedRating struct has nil ID")
	}
	if _, exists := d.ratings[updatedRating.ID]; !exists {
		log.Error("updatedRating ID does not exists in mock db")
		return errors.New("updatedRating ID does not exists in mock db")
	}
	d.ratings[updatedRating.ID] = updatedRating
	return nil
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
