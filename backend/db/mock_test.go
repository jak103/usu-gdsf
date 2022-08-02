package db

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/google/uuid"
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
	id := keys[rand.Intn(len(keys))].Interface().(uuid.UUID)
	game, _ := mock.GetGameByID(id)

	assert.Equal(t, game.ID, id)
}
