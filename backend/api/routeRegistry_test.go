package api

import (
	"net/http"
	"reflect"
	"runtime"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRouteRegistry(t *testing.T) {
	myTestRoute := route{method: http.MethodGet, path: "/test/path", handler: testHandler}
	registerRoute(myTestRoute)

	assert.GreaterOrEqual(t, len(routes), 1)

	found := false
	for _, route := range routes {
		if route.method == myTestRoute.method &&
			route.path == myTestRoute.path &&
			runtime.FuncForPC(reflect.ValueOf(route.handler).Pointer()).Name() ==
				runtime.FuncForPC(reflect.ValueOf(myTestRoute.handler).Pointer()).Name() {
			found = true
		}
	}

	assert.True(t, found)
}

func testHandler(c echo.Context) error {
	// Test handler does nothing
	return nil
}
