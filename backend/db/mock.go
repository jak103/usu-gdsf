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
