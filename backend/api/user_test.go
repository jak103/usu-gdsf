package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jak103/usu-gdsf/auth"
	"github.com/stretchr/testify/assert"
)

func AssertResponseCode(t *testing.T, method string, path string, expectedCode int) bool {
	e := echo.New()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/game", nil)
	c := e.NewContext(request, recorder)

	if assert.NoError(t, getAllGames(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
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
