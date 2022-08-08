package api

import "github.com/labstack/echo/v4"

type route struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

var routes []route = make([]route, 0)
var restrictedRoutes []route = make([]route, 0)

// This function adds a route anyone can access, no JWT necessary
func registerRoute(r route) {
	routes = append(routes, r)
}

// This function adds a restricted route (one that requires a JWT to access)
func registerRestrictedRoute(r route) {
	restrictedRoutes = append(restrictedRoutes, r)
}
