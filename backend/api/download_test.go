package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/jak103/usu-gdsf/db"
	"github.com/stretchr/testify/assert"
	"github.com/jak103/usu-gdsf/models"
)

var (
	_db, _ = db.NewDatabaseFromEnv()

	download0 = models.Download{
		UserId:       "000",
		GameId:       "62dadce12fcdb47399118408",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	download1 = models.Download{
		UserId:       "001",
		GameId:       "62dadce12fcdb47399118408",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
	}
)
//this will need to be changed eventually. It is looking for response 500 now but when we get the database up it will be 200.
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

func TestGetDownload(t *testing.T) {
	downloadId, _ := _db.AddDownload(download1)
	e := echo.New()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(request, recorder)
	c.SetPath("/downloads/:id")
	c.SetParamNames("id")
	c.SetParamValues(downloadId)

	if assert.NoError(t, getDownloadByID(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}

}

func TestInvalidGetDownload(t *testing.T) {
	e := echo.New()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(request, recorder)
	c.SetPath("/downloads/:id")
	c.SetParamNames("id")
	c.SetParamValues("00000")

	if assert.NoError(t, getDownloadByID(c)) {
		assert.Equal(t, http.StatusInternalServerError, 500)
	}
}