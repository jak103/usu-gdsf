package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

//this will need to be changed eventually. It is looking for response 500 now but when we get the database up it will be 200.
func TestGetAllGames(t *testing.T) {
	e := echo.New()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/game", nil)
	c := e.NewContext(request, recorder)

	if assert.NoError(t, getAllGames(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}
