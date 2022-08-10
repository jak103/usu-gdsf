package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllDownloads(t *testing.T) {
	e := echo.New()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/downloads", nil)
	c := e.NewContext(request, recorder)

	if assert.NoError(t, getAllDownloads(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}

func TestAddGame(t *testing.T) {
	_, err := _db.AddDownload(download0)
	assert.Nil(t, err)
}
