package api

import (
	"net/http"

	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

func user(c echo.Context) error {
	return c.JSON(http.StatusOK, "User get handler")
}

func createUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "Created User")
}

func init() {
	log.Info("Running user init")

	registerRoute(route{method: http.MethodGet, path: "/user", handler: user})
	registerRoute(route{method: http.MethodPost, path: "/user", handler: createUser})
}
