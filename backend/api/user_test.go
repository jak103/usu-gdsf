package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/jak103/usu-gdsf/auth"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

func AssertResponseCode(t *testing.T, method string, path string, expectedCode int) bool {
	e := echo.New()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, nil)
	c := e.NewContext(request, recorder)

	if assert.NoError(t, getAllGames(c)) {
		return expectedCode == recorder.Code
	} else {
		return false
	}
}

func TestUserRoute(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/user", 200))
}

func TestRegisterRoute(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/user/register", 200))
}

func TestDownloadsRoute(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/user/downloads", 200))
}

func TestPasswordHashing(t *testing.T, testpass string) {
	var s Server = *NewServer(&sync.WaitGroup{})
	s.Start()

}
