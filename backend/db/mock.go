package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

var _ Database = (*Mock)(nil)

type Mock struct {
	games map[uuid.UUID]models.Game
}

func (db *Mock) GetGameByID(id uuid.UUID) (*models.Game, error) {
	if id.String() != "" {
		bs, _ := json.Marshal(db.games)
		fmt.Println(string(bs))
		if game, ok := db.games[id]; ok {
			return &game, nil
		}
	}
	return nil, errors.New("mockdb: game not found")
}

func (db *Mock) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)

	for _, game := range db.games {
		games = append(games, game)
	}

	return games, nil
}

func (d *Mock) GetGamesByTags(tags []string) ([]models.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) GetGamesByPublishDate(startRange string, endRange string) ([]models.Game, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) CreateGame(newGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) DeleteGame(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) UpdateGame(updatedGame models.Game) error {
	panic("not implemented") // TODO: Implement
}

// Users
func (d *Mock) GetAllUsers() ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) GetUserByID(id uuid.UUID) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) GetUsersByRole(role int64) ([]models.User, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) CreateUser(newUser models.User) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) DeleteUser(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) UpdateUser(updatedUser models.User) error {
	panic("not implemented") // TODO: Implement
}

// Ratings
func (d *Mock) GetRatingByID(id uuid.UUID) (*models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) GetRatingsByGame(gameID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) GetRatingsByUser(userID uuid.UUID) ([]models.GameRating, error) {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) CreateRating(newRating models.GameRating) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) DeleteRating(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) DeleteRatingsByGame(gameID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (d *Mock) updateRating(updatedRating models.GameRating) error {
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

	log.Debug("Sample Database Initialized")

	return nil
}

func (db Mock) Disconnect() error {
	return nil
}

// func init(){

// }
