package api

import (
	"archive/zip"
	"bytes"
	b64 "encoding/base64"
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/gcs"
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

func gameDataDownload(c echo.Context) error {
	downloadType := c.QueryParam("downloadType")

	if downloadType != "screenshots" && downloadType != "game" {
		return echo.NewHTTPError(http.StatusBadRequest, "download type must be of type 'Screenshots' or 'Game'")
	}

	bucket := c.QueryParam("bucket")
	if bucket == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bucket name not set correctly")
	}

	objectManager := gcs.NewGcsObjectManager(bucket)

	objects, err := objectManager.ListObjects(downloadType)
	if err != nil {
		log.WithError(err).Error("Error getting the list of objects")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	downloadManager := gcs.NewGcsDataDownloader()

	for idx, object := range objects {
		objectData, err := downloadManager.DownloadData(bucket, object.ObjectName)
		if err != nil {
			log.WithError(err).Error("An error occurred while fetching the data")
			return err
		}

		objects[idx].SetData(b64.StdEncoding.EncodeToString(objectData))
	}

	if downloadType == "screenshots" {
		return c.JSON(http.StatusOK, []interface{}{objects})
	}

	buf := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buf)

	for _, object := range objects {
		zipFile, err := zipWriter.Create(object.ObjectName)
		if err != nil {
			log.WithError(err).Error("An error occurred while creating the zip file")
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		bytesDecoded, err := b64.StdEncoding.DecodeString(object.ObjectData)

		if err != nil {
			log.WithError(err).Error("Unable to convert base64 to byte array")
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to parse base64 string to byte array")
		}

		_, err = zipFile.Write(bytesDecoded)
		if err != nil {
			log.WithError(err).Error("Unable to write bytes to zip file")
			return echo.NewHTTPError(http.StatusInternalServerError, "Unable to write bytes to zip file")
		}
	}

	if err := zipWriter.Close(); err != nil {
		log.WithError(err).Error("Unable to close the zip file")
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to close the zip file")
	}

	return c.Blob(http.StatusOK, "application/zip", buf.Bytes())
}

func getGames(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()
	query := c.Request().URL.Query()

	if query.Has("userid") {
		userids := query["userid"]
		// db.GetGamesByUserId(userids[0])
		return c.JSON(http.StatusOK, "List of games by "+userids[0])
	}

	if query.Has("tag") {
		tags := query["tag"]
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
	registerRoute(route{method: http.MethodGet, path: "/game/download", handler: gameDataDownload})
	registerRoute(route{method: http.MethodGet, path: "/game/:id", handler: getGameByID})
	registerRoute(route{method: http.MethodGet, path: "/game", handler: getGames})
	registerRoute(route{method: http.MethodPost, path: "/game/highest-rated", handler: getHighestRatedGames})
	registerRestrictedRoute(route{method: http.MethodPost, path: "/game/add", handler: newGameHandler})
}
