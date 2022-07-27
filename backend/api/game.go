package api

import (
	"github.com/jak103/usu-gdsf/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

const (
	// TODO these are placeholder form names for adding a new game ('/game' POST)
	// TODO match these strings with the view's form names for '/game' POST
	NAME    = "Name"
	AUTHOR  = "Author"
	VERSION = "Version"
)

func game(c echo.Context) error {
	return c.JSON(http.StatusOK, "Successful game get!")
}

func getAllGames(c echo.Context) error {
	_db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	if result, err := _db.GetAllGames(); err != nil {
		log.Error("An error occurred while getting game records: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{result})
	}
}

func newGameHandler(c echo.Context) error {
	// create random uint64 for new game ID
	var id = rand.Uint64()
	for {
		// check for already existing ID
		exists, _ := db.IdExists(id)
		if !exists {
			break
		} else {
			id = rand.Uint64()
		}
	}

	// create new game db model
	newGame := models.Game{
		Id:           id,
		Name:         c.FormValue(NAME),
		Author:       c.FormValue(AUTHOR),
		CreationDate: time.Now().Format("MM/DD/YYYY"),
		Version:      c.FormValue(VERSION),
	}

	// Add new game to database
	_db, dbGetErr := db.NewDatabaseFromEnv()
	err := _db.AddGame(newGame)
	// TODO error handling

	// TODO register new route with ID

	return c.JSON(http.StatusOK, "New game handler")
}

func init() {
	log.Info("Running game init")
	registerRoute(route{method: http.MethodGet, path: "/game", handler: game})
	registerRoute(route{method: http.MethodGet, path: "/games", handler: getAllGames})
	registerRoute(route{method: http.MethodPost, path: "/game", handler: newGameHandler})
}
