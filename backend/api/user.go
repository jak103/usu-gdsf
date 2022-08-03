package api

import (
	"net/http"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

func user(c echo.Context) error {
	return c.JSON(http.StatusOK, "User get handler")
}

func createUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Created User")
}

func returnAllUsers(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	// TODO make sure this is checked when the DB call is properly implemented.
	if result, err := db.GetAllUsers(); err != nil {
		log.Error("An error occurred while getting a record of all users: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{result})
	}
}

func init() {
	log.Info("Running user init")

	registerRoute(route{method: http.MethodGet, path: "/login", handler: user})
	registerRoute(route{method: http.MethodPost, path: "/register", handler: createUser})
	registerRoute(route{method: http.MethodGet, path: "/allusers", handler: returnAllUsers})
}
