package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
)

func getGameByID(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}
	gameID, _ := uuid.Parse(c.Param("id"))

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
	db, err := db.NewDatabaseFromEnv()
	query := c.Request().URL.Query()

	if query.Has("userid") {
		userids, _ := query["userid"]
		// db.GetGamesByUserId(userids[0])
		return c.JSON(http.StatusOK, "List of games by "+userids[0])
	}

	if query.Has("tag") {
		tags, _ := query["tag"]
		if result, err := db.GetGamesByTags(tags); err != nil {
			log.Error("An error occurred while getting game records: %v", err)
			return err
		} else {
			return c.JSON(http.StatusOK, []interface{}{result})
		}

	}

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
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	g := *new(models.Game)
	if err = c.Bind(&g); err != nil {
		return err
	}
	g.SetUUID()

	if err := db.CreateGame(g); err != nil {
		log.Error("An error occurred while creating user records %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, "New Game Added")
	}
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
	registerRestrictedRoute(route{method: http.MethodPost, path: "/game/add", handler: newGameHandler})
}
