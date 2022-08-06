package db

import (
	"strconv"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

var _ Database = (*Mock)(nil)

type Mock struct {
	games map[string]models.Game
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
