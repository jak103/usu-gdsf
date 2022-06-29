package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}
