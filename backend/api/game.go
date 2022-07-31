package api

import (
	"fmt"
	"net/http"
	"strings"

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

// for downloaded games, this route allows you to play them
func gameDownload(c echo.Context) error {
	return c.JSON(http.StatusOK, "Download a game")
}

func getGames(c echo.Context) error {
	query := c.Request().URL.Query()

	if query.Has("userid") {
		userids, _ := query["userid"]
		// db.GetGamesByUserId(userids[0])
		return c.JSON(http.StatusOK, "List of games by " + userids[0]);
	}

	if query.Has("tag") {
		tags, _ := query["tag"]
		// db.GetGamesByTags(tags)
		return c.JSON(http.StatusOK, fmt.Sprintf("List of games with tags %s", strings.Join(tags, ", ")))
	}

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
	//registerRoute(route{method: http.MethodGet, path: "/game", handler: game})
	//registerRoute(route{method: http.MethodGet, path: "/game/download", handler: gameDownload})
	//registerRoute(route{method: http.MethodGet, path: "/games", handler: getGames})
	//registerRoute(route{method: http.MethodPost, path: "/game", handler: newGameHandler})
  registerRoute(route{method: http.MethodGet, path: "/game/download", handler: gameDownload})
	registerRoute(route{method: http.MethodGet, path: "/game/:id", handler: getGameByID})
	registerRoute(route{method: http.MethodGet, path: "/game", handler: getGames})
	registerRoute(route{method: http.MethodPost, path: "/game/add", handler: newGameHandler})
}
