package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
)

const (
	// TODO these are placeholder form var names for adding a new game
	// TODO match these strings with the view's form var names for '/game' POST
	NAME      = "Name"
	DEVELOPER = "Developer"
	VERSION   = "Version"
	LINK      = "DownloadLink"
)

var v = validator.New()

func getGame(c echo.Context) error {

	// get id from path
	_id := c.Param("id")

	// get game from db with id
	_db, getDbErr := db.NewDatabaseFromEnv()
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	game, err := _db.GetGameByID(_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database find game ID error")
	}
	return c.JSON(http.StatusOK, game)
}

func getGamesWithTags(c echo.Context) error {
	// connect the db
	_db, err := db.NewDatabaseFromEnv()
	if err != nil {
		log.WithError(err).Error("Database connection error in API getGamesWithTags")
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}

	// grab tags from context
	rawTags := c.QueryParam("tags")
	// split string into array
	tags := strings.Split(rawTags, "-")
	// fetch games with tags
	games, err := _db.GetGamesByTags(tags, false)

	if err != nil {
		log.WithError(err).Error("Database GetGamesByTags error in API getGamesWithTags")
		return c.JSON(http.StatusInternalServerError, "Database fetch games with tags error")
	}

	return c.JSON(http.StatusOK, games)
}

func getAllGames(c echo.Context) error {
	_db, err := db.NewDatabaseFromEnv()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")

	}

	if result, err := _db.GetAllGames(); err != nil {
		return c.JSON(http.StatusInternalServerError, "error in GetAllGameS API")

	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func sortAllGame(c echo.Context) error {

	_db, err := db.NewDatabaseFromEnv()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}

	// srt=name-desc
	sortConfig := c.QueryParam("srt")
	sortConfigArray := strings.Split(sortConfig, "-")

	if len(sortConfigArray) != 2 {
		return c.JSON(http.StatusInternalServerError, "not enough parameter to address sorting job")
	}

	if (sortConfigArray[1] == "ASC") || (sortConfigArray[1] == "asc") || (sortConfigArray[1] == "DSC") || (sortConfigArray[1] == "dsc") {
		order := 1

		if sortConfigArray[1] == "asc" || sortConfigArray[1] != "ASC" {
			order = -1
		}

		games, err := _db.SortGames(sortConfigArray[0], order)

		if err != nil {
			log.WithError(err).Error("Sort data error in API SortGame")
			return c.JSON(http.StatusInternalServerError, "couldn't get sorted data because of server error")
		}

		return c.JSON(http.StatusOK, games)

	}

	return c.JSON(http.StatusInternalServerError, "not a valid order of sorting")

}

func getGamesWithFirstLetter(c echo.Context) error {
	// connect the db
	_db, err := db.NewDatabaseFromEnv()
	if err != nil {
		log.WithError(err).Error("Database connection error in API getGamesWithFirstLetter")
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}

	// grab tags from context
	letter := c.QueryParam("ltr")
	// split string into array    // fetch games with tags
	games, err := _db.GetGamesByFirstLetter(letter)

	if err != nil {
		log.WithError(err).Error("Database function GetGamesByFirstLetter error")
		return c.JSON(http.StatusInternalServerError, "Database fetch games with first letter error")
	}

	return c.JSON(http.StatusOK, games)
}

func newGameHandler(c echo.Context) error {
	// create new game model
	// TODO need a security layer in between the form and our new game struct
	newGame := models.Game{
		Name:         c.FormValue(NAME),
		Developer:    c.FormValue(DEVELOPER),
		CreationDate: time.Now(),
		Version:      c.FormValue(VERSION),
		DownloadLink: c.FormValue(LINK),
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
	registerRoute(route{method: http.MethodGet, path: "/game/:id", handler: getGame})
	registerRoute(route{method: http.MethodGet, path: "/game/tags", handler: getGamesWithTags})
	registerRoute(route{method: http.MethodGet, path: "/games/sort", handler: sortAllGame})
	registerRoute(route{method: http.MethodGet, path: "/games/firstLetter", handler: getGamesWithFirstLetter})
}
