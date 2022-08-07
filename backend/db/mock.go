package db

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

var _ Database = (*Mock)(nil)

type Mock struct {
	games     map[string]models.Game
	downloads map[string]models.Download
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

// GetGamesByTags search and return all games with given tags
func (db Mock) GetGamesByTags(tags []string, matchAll bool) ([]models.Game, error) {
	games := make([]models.Game, 0)
	for _, game := range db.games {
		if matchAll {
			matched := []string{}
			for _, t := range game.Tags {
				if containsTag(tags, t) {
					matched = append(matched, t)
				}
			}
			if len(matched) == len(tags) {
				games = append(games, game)
			}
		} else {
			for _, t := range game.Tags {
				if containsTag(tags, t) {
					games = append(games, game)
					break
				}
			}
		}

	}
	return games, nil
}

// Helper function to check if one array contains an element
func containsTag(a []string, el string) bool {
	for _, v := range a {
		if v == el {
			return true
		}
	}
	return false
}

func (db Mock) GetGameByID(id string) (models.Game, error) {
	return db.games[id], nil
}

func (db Mock) AddGame(game models.Game) (string, error) {
	var id = strconv.Itoa(len(db.games) + 1)
	db.games[id] = game
	return id, nil
}

func (db Mock) AddDownload(download models.Download) (string, error) {
	var id = strconv.Itoa(len(db.downloads) + 1)
	db.downloads[id] = download
	return id, nil
}

func (db Mock) GetAllGames() ([]models.Game, error) {
	games := make([]models.Game, 0)

	for _, game := range db.games {
		games = append(games, game)
	}

	return games, nil
}

func (db Mock) GetAllDownloads() ([]models.Download, error) {
	downloads := make([]models.Download, 0)

	for _, download := range db.downloads {
		downloads = append(downloads, download)
	}

	return downloads, nil
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
