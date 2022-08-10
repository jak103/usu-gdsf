package api

import (
	"net/http"
	"sort"

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

type gameWithRating struct {
	game   models.Game
	rating float32
}

func getAvgGameRatings() ([]gameWithRating, error) {
	db, err := db.NewDatabaseFromEnv()
	gamesWithRatings := []gameWithRating{}

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return nil, err
	}

	if games, err := db.GetAllGames(); err != nil {
		for _, game := range games {
			avg := 0.0
			ratings, err := db.GetRatingsByGame(game.ID)
			if err != nil {
				log.WithError(err).Error("Unable to get ratings")
				return nil, err
			}

			for _, rating := range ratings {
				avg += float64(rating.RatingValue)
			}
			avg = avg / float64(len(ratings))
			gamesWithRatings = append(gamesWithRatings, gameWithRating{game: game, rating: float32(avg)})
		}
	}

	return gamesWithRatings, nil
}

func getHighestRatedGames(c echo.Context) error {
	gamesWithRatings, err := getAvgGameRatings()

	if err != nil {
		log.WithError(err).Error("Unable to get games with ratings")
		return err
	} else {
		sort.SliceStable(gamesWithRatings, func(i, j int) bool {
			return gamesWithRatings[i].rating > gamesWithRatings[j].rating
		})
	}

	return c.JSON(http.StatusOK, gamesWithRatings)
}

func init() {
	log.Info("Running game init")
	registerRoute(route{method: http.MethodGet, path: "/game/download", handler: gameDownload})
	registerRoute(route{method: http.MethodGet, path: "/game/:id", handler: getGameByID})
	registerRoute(route{method: http.MethodGet, path: "/game", handler: getGames})
	registerRoute(route{method: http.MethodPost, path: "/game/highest-rated", handler: getHighestRatedGames})
	registerRestrictedRoute(route{method: http.MethodPost, path: "/game/add", handler: newGameHandler})
}
