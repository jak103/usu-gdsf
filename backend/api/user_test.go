package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jak103/usu-gdsf/auth"
	"github.com/stretchr/testify/assert"
)

func AssertResponseCode(t *testing.T, method string, path string, expectedCode int) bool {
	params := auth.TokenParams{
		Type:      auth.ACCESS_TOKEN,
		UserId:    42,
		UserEmail: "tst@example.com",
	}

	token, _ := auth.GenerateToken(params)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, nil)
	request.AddCookie(&http.Cookie{
		Name:  "accessToken",
		Value: token,
	})
	GlobalTestServer.echo.ServeHTTP(recorder, request)

	return expectedCode == recorder.Code
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
