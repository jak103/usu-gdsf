package api

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func robots(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Robots!")
}

func init() {
	registerRoute(route{method: http.MethodGet, path: "/robots.txt", handler: robots})
}