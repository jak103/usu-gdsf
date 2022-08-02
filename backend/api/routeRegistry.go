package api

import "github.com/labstack/echo/v4"

type route struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

var routes []route = make([]route, 0)
var restrictedRoutes []route = make([]route, 0)

func registerRoute(r route) {
	routes = append(routes, r)
}

func registerRestrictedRoute(r route) {
	restrictedRoutes = append(restrictedRoutes, r)
}
