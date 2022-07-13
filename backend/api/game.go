package api

import (
	"net/http"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

func game(c echo.Context) error {
	return c.JSON(http.StatusOK, "Successful game get!")
}

func gameRecords(c echo.Context) error {
	db, err := db.GetDb()
	if err != nil {
		return err
	}

	log.Debug("Database Name: %v", db.Name)

	if result, err := db.GetAllGameRecords(); err != nil {
		log.Error("An error occurred while getting game records: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{ result })
	}
}

func init() {
	registerRoute(route{method: http.MethodGet, path: "/game", handler: game})
	registerRoute(route{method: http.MethodGet, path: "/game/records", handler: gameRecords})
}
