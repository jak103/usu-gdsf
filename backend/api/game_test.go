package api

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s Server = *NewServer(&sync.WaitGroup{})

func init() {
	s.Start()
}

func AssertResponseCode(t *testing.T, method string, path string, expectedCode int) bool {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, nil)
	s.echo.ServeHTTP(recorder, request)
	return expectedCode == recorder.Code
}

func TestGame(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/game", 200))
}

//this will need to be changed eventually. It is looking for response 500 now but when we get the database up it will be 200.
func TestGetAllGames(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/games", 500))
}
