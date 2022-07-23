package api

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertResponseCode(t *testing.T, method string, path string, expectedCode int) bool {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, nil)
	S.echo.ServeHTTP(recorder, request)
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
