package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameRoute(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/game", 200))
}

//this will need to be changed eventually. It is looking for response 500 now but when we get the database up it will be 200.
func TestGetAllGames(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/games", 500))
}
