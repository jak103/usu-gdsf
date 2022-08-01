package db

import (
	"errors"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"reflect"
	"strconv"
)

var _ Database = (*Mock)(nil)

type Mock struct {
	games map[string]models.Game
}

// RemoveGame removes the given game from the db
func (db Mock) RemoveGame(game models.Game) error {
	deleteKey := ""
	for i, v := range db.games {
		if reflect.DeepEqual(v, game) {
			deleteKey = i
			break
		}
	}

	if deleteKey == "" {
		return errors.New("mockdb: can't find game in RemoveGame")
	}

	delete(db.games, deleteKey)
	return nil
}

// GetGamesByTag search and return all games with given tag
func (db Mock) GetGamesByTag(tag string) ([]models.Game, error) {
	games := make([]models.Game, 0)
	for _, game := range db.games {
		for _, t := range game.Tags {
			if t == tag {
				games = append(games, game)
			}
		}
	}
	return games, nil
}

func (db Mock) GetGameID(game models.Game) (string, error) {
	for i, v := range db.games {
		if reflect.DeepEqual(v, game) {
			return i, nil
		}
	}
	return "", errors.New("mockDB: ID couldn't be found with given game")
}

func (db Mock) GetGameByID(id string) (models.Game, error) {
	return db.games[id], nil
}

func (db Mock) AddGame(game models.Game) (string, error) {
	var id = strconv.Itoa(len(db.games) + 1)
	db.games[id] = game
	return id, nil
}

func (db Mock) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)

	for _, game := range db.games {
		games = append(games, game)
	}

	return games, nil
}

func (db Mock) Connect() error {
	log.Info("mock connect")

	if len(db.games) == 0 {
		for i, v := range CreateGamesFromJson() {
			db.games[strconv.Itoa(i)] = v
		}
	}

	log.Debug("Sample Database Initialized")

	return nil
}

func (db Mock) Disconnect() error {
	return nil
}
