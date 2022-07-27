package api

import (
	"net/http"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

func getGameByID(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}
	gameID := c.Param("id")

	if result, err := db.GetGameByID(gameID); err != nil {
		log.Error("An error occurred while getting game records: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{result})
	}

}

func getAllGames(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	if result, err := db.GetAllGames(); err != nil {
		log.Error("An error occurred while getting game records: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{result})
	}
}

func newGameHandler(c echo.Context) error {
	// TODO #5 We should probably actually create a game here

	return c.JSON(http.StatusOK, "New game handler")
}

func init() {
	log.Info("Running game init")
	registerRoute(route{method: http.MethodGet, path: "/game/:id", handler: getGameByID})
	registerRoute(route{method: http.MethodGet, path: "/game", handler: getAllGames})
	registerRoute(route{method: http.MethodPost, path: "/game/add", handler: newGameHandler})
}
