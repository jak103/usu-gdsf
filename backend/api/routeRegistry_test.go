package api

import (
	"net/http"
	"reflect"
	"runtime"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func findTestRouter(r *route, routes []route) bool {
	for _, route := range routes {
		if route.method == r.method &&
			route.path == r.path &&
			runtime.FuncForPC(reflect.ValueOf(route.handler).Pointer()).Name() ==
				runtime.FuncForPC(reflect.ValueOf(r.handler).Pointer()).Name() {
			return true
		}
	}

	return false

}

func TestRouteRegistry(t *testing.T) {
	myTestRouter := route{method: http.MethodGet, path: "/test/path", handler: testHandler}
	registerRoute(myTestRouter)

	assert.True(t, findTestRouter(&myTestRouter, routes))

	myRestrictedTestRoute := route{method: http.MethodGet, path: "/test/restricted", handler: testHandler}
	registerRestrictedRoute(myRestrictedTestRoute)

	assert.True(t, findTestRouter(&myRestrictedTestRoute, restrictedRoutes))

	assert.GreaterOrEqual(t, len(routes), 1)

}

func testHandler(c echo.Context) error {
	// Test handler does nothing
	return nil
}
