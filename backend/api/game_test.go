package api

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/models"
	"github.com/stretchr/testify/assert"
)

// func AssertResponseCode(t *testing.T, method string, path string, expectedCode int) bool {
// 	recorder := httptest.NewRecorder()
// 	request := httptest.NewRequest(method, path, nil)
// 	s.echo.ServeHTTP(recorder, request)
// 	return expectedCode == recorder.Code
// }

func TestGame(t *testing.T) {
	var s Server = *NewServer(&sync.WaitGroup{})
	s.Start()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/game", nil)
	s.echo.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

//this will need to be changed eventually. It is looking for response 500 now but when we get the database up it will be 200.
func TestGetAllGames(t *testing.T) {
	var s Server = *NewServer(&sync.WaitGroup{})
	s.Start()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/game", nil)
	s.echo.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	// assert.True(t, AssertResponseCode(t, http.MethodGet, "/games", 500))
}

func TestGameInfoHandler(t *testing.T) {
	game := models.Game{
		Name:    "Test Game",
		Developer:  "John Doe",
		Version: "1.0",
	}

	_db, _ := db.NewDatabaseFromEnv()
	id, _ := _db.AddGame(game)

	var s Server = *NewServer(&sync.WaitGroup{})
	s.Start()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/game/" + id, nil)
	s.echo.ServeHTTP(recorder, request)

	t.Log(game, _db, "Hello, World ----------------------------------------")
	t.Log(recorder)
	t.Log("Hello, World ----------------------------------------")

	assert.Equal(t, http.StatusOK, recorder.Code)
}
