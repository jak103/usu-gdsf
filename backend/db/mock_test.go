package db

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	mock := Mock{}
	mock.Connect()
}

func TestGetGameByID(t *testing.T) {
	mock := Mock{}
	mock.Connect()

	keys := reflect.ValueOf(mock.games).MapKeys()
	id := keys[rand.Intn(len(keys))].Interface().(string)
	game, _ := mock.GetGameByID(id)

	assert.Equal(t, game.ID.String(), id)
}

func TestGetAllGames(t *testing.T) {
	mock := Mock{}
	mock.Connect()
	allGames, _ := mock.GetAllGames()
	assert.Equal(t, len(mock.games), len(allGames))
	assert.Greater(t, len(allGames), 0)
}
