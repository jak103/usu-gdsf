package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jak103/usu-gdsf/auth"
	"github.com/stretchr/testify/assert"
)

// func AssertResponseCode(t *testing.T, method string, path string, expectedCode int) bool {
// 	recorder := httptest.NewRecorder()
// 	request := httptest.NewRequest(method, path, nil)
// 	s.echo.ServeHTTP(recorder, request)
// 	return expectedCode == recorder.Code
// }

func TestGame(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/game", nil)
	GlobalTestServer.echo.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

//this will need to be changed eventually. It is looking for response 500 now but when we get the database up it will be 200.
func TestGetAllGames(t *testing.T) {
	params := auth.TokenParams{
		Type:      auth.ACCESS_TOKEN,
		UserId:    42,
		UserEmail: "tst@example.com",
	}

	token, _ := auth.GenerateToken(params)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/game", nil)
	request.Header.Set("accessToken", token)

	GlobalTestServer.echo.ServeHTTP(recorder, request)
	println("this that")
	println(recorder.Body)
	assert.Equal(t, http.StatusOK, recorder.Code)
	// assert.True(t, AssertResponseCode(t, http.MethodGet, "/games", 500))
}
