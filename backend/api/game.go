package api

import (
	"fmt"
	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	// TODO these are placeholder form var names for adding a new game
	// TODO match these strings with the view's form var names for '/game' POST
	NAME    = "Name"
	AUTHOR  = "Author"
	VERSION = "Version"
)

func game(c echo.Context) error {
	// get id from path
	id := c.Path()[6:]

	// get game from db with id
	_db, getDbErr := db.NewDatabaseFromEnv()
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	game, err := _db.GetGameByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database find game ID error")
	}

	// TODO return view filled by game details
	return c.JSON(http.StatusOK, fmt.Sprintf("%s\n%s\n%s\n%s", game.Name, game.Author, game.CreationDate, game.Version))
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
	// create new game model
	// TODO need a security layer in between the form and our new game struct
	newGame := models.Game{
		Name:         c.FormValue(NAME),
		Author:       c.FormValue(AUTHOR),
		CreationDate: time.Now().Format("MM/DD/YYYY"),
		Version:      c.FormValue(VERSION),
	}

	// Add new game to database
	_db, getDbErr := db.NewDatabaseFromEnv()
	id, err := _db.AddGame(newGame)

	// error handling
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database add game error")
	}

	// register new route with ID
	registerRoute(route{method: http.MethodGet, path: fmt.Sprintf("/game/%s", id), handler: game})

	// TODO return successful game add data
	return c.JSON(http.StatusOK, "New game handler")
}

func init() {
	log.Info("Running game init")
	registerRoute(route{method: http.MethodGet, path: "/games", handler: getAllGames})
	registerRoute(route{method: http.MethodPost, path: "/game", handler: newGameHandler})

	// register routes for all games from the db
	log.Info("Creating routes for all games in database")
	_db, getDbErr := db.NewDatabaseFromEnv()
	if getDbErr != nil {
		return
	}
	games, getGamesErr := _db.GetAllGames()
	if getGamesErr != nil {
		return
	}
	for _, v := range games {
		gameID, getIdErr := _db.GetGameID(v)
		if getIdErr != nil {
			continue
		}
		registerRoute(route{method: http.MethodGet, path: fmt.Sprintf("/info/%s", gameID), handler: game})
	}
}
