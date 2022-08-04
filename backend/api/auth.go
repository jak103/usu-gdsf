package api

import (
	"net/http"

	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
)

func helloWorld(c echo.Context) error {
	log.Info("Running hello world handler!")
	return c.JSON(http.StatusOK, "Hellow orld!")
}

func init() {
	registerRoute(route{
		method: http.MethodGet,
		path: "/test",
		handler: helloWorld,
		requireAuth: true,
	})
}
