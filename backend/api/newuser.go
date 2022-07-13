package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func newuser(c echo.Context) error {
	return c.JSON(http.StatusOK, "NewUser Gets a New Handler")
}

func init() {
	registerRoute(route{method: http.MethodGet, path: "/newuser", handler: newuser})
}
