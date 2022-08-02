package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

var _ Database = (*Mock)(nil)

type Mock struct {
	games map[string]models.Game
}

func (db *Mock) GetGameByID(id string) (*models.Game, error) {
	if id != "" {
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

func (db *Mock) Connect() error {
	log.Info("mock connect")

	if len(db.games) == 0 {
		games := make(map[string]models.Game)
		for _, v := range CreateGamesFromJson() {
			games[v.ID.String()] = v
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
