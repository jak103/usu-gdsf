package api

import (
	"net/http"
	"time"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
)

const (
	USER      = "UserId"
	GAME      = "GameId"
)

func getAllDownloads(c echo.Context) error {
	_db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	if result, err := _db.GetAllDownloads(); err != nil {
		log.Error("An error occurred while getting download records: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{result})
	}
}

func getDownloadByID(c echo.Context) error {
	id := c.Param("id")

	_db, getDbErr := db.NewDatabaseFromEnv()
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	download, err := _db.GetDownloadByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Database find download ID error")
	}

	return c.JSON(http.StatusOK, download)
}

// Method used to handle creating a new Download Record
func newDownloadHandler( c echo.Context) error {
	newDownload := models.Download{
		UserId: c.FormValue(USER),
		GameId: c.FormValue(GAME),
		CreationDate: time.Now(),
	}

	_db, getDbErr := db.NewDatabaseFromEnv()
	_, err := _db.AddDownload(newDownload)

	// error handling
	if getDbErr != nil {
		return c.JSON(http.StatusInternalServerError, "Database connection error")
	}
	if err != nil {
		log.Error("An error occurred while getting download records: %v", err)
		return c.JSON(http.StatusInternalServerError, "Database add download error")
	}

	message := "Download added"
	return c.JSON(http.StatusOK, message)
}

func init() {
	registerRoute(route{method: http.MethodGet,  path: "/downloads", handler: getAllDownloads})
	registerRoute(route{method: http.MethodGet,  path: "/downloads/:id", handler: getDownloadByID})
	registerRoute(route{method: http.MethodPost, path: "/downloads", handler: newDownloadHandler})
	
}