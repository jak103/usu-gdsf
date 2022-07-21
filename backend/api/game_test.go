package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameRoute(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/game", 200))
}

func TestGetAllGames(t *testing.T) {
	assert.True(t, AssertResponseCode(t, http.MethodGet, "/games", 500) )
}
