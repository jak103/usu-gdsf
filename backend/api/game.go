package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func game(c echo.Context) error {
	return c.JSON(http.StatusOK, "Successful game get!")
}

func init() {
	registerRoute(route{method: http.MethodGet, path: "/game", handler: game})
}
