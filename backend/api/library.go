package api

import (
	"net/http"
	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

func library(c echo.Context) error {
	return c.JSON(http.StatusOK, "Library handler")
}

func init() {
	log.Info("Running libary init")
	registerRoute(route{method: http.MethodGet, path: "/library", handler: library})
}
