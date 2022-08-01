package api

import "github.com/labstack/echo/v4"

type route struct {
	method      string
	path        string
	handler     echo.HandlerFunc
	requireAuth bool
}

var routes []route = make([]route, 0)

func registerRoute(r route) {
	routes = append(routes, r)
}
