package api

import (
	"net/http"

	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

func user(c echo.Context) error {
	return c.JSON(http.StatusOK, "User get handler")
}

func register(c echo.Context) error {
	// User registration screen
	return c.JSON(http.StatusOK, "User registration handler")
}

func downloads(c echo.Context) error {
	// Return games that a user has downloaded/played
	return c.JSON(http.StatusOK, "User downloads handler")
}

func init() {
	log.Info("Running user init")

	registerRoute(route{method: http.MethodGet, path: "/user", handler: user})
	registerRoute(route{method: http.MethodGet, path: "/user/register", handler: register})
	registerRoute(route{method: http.MethodGet, path: "user/downloads", handler: downloads})
}
