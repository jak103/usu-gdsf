package api

import (
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

func gameInfoHandler(c echo.Context) error {
	// get id from path
	id := c.Param("id")

	// get game from db with id
	_db, getDbErr := db.NewDatabaseFromEnv()
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	game, err := _db.GetGameByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database find game ID error")
	}

	return c.JSON(http.StatusOK, game)
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
		CreationDate: time.Now(),
		Version:      c.FormValue(VERSION),
	}

	// Add new game to database
	_db, getDbErr := db.NewDatabaseFromEnv()
	_, err := _db.AddGame(newGame)

	// error handling
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database add game error")
	}

	// TODO return successful game add
	return c.JSON(http.StatusOK, "New game handler")
}

func init() {
	log.Info("Running game init")
	registerRoute(route{method: http.MethodGet, path: "/games", handler: getAllGames})
	registerRoute(route{method: http.MethodPost, path: "/game", handler: newGameHandler})
	registerRoute(route{method: http.MethodGet, path: "/info/:id", handler: gameInfoHandler})
}
