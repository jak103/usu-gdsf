package db

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/models"
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

func TestGetAllGames(t *testing.T) {
	mock := Mock{}
	mock.Connect()
	allGames, _ := mock.GetAllGames()
	assert.Equal(t, len(mock.games), len(allGames))
	assert.Greater(t, len(allGames), 0)
}

func TestCreateGame(t *testing.T) {
	mock := Mock{}
	mock.games = make(map[uuid.UUID]models.Game)
	gameID := uuid.New()
	userID := uuid.New()
	newGame := models.Game{
		ID:               gameID,
		Title:            "NewGame",
		Description:      "lame",
		UserID:           userID,
		VersionNumber:    "1.0.0",
		PublishTimestamp: "7/29/2022",
	}
	err := mock.CreateGame(newGame)
	assert.Contains(t, mock.games, gameID)
	assert.Equal(t, nil, err)
	// verify that you cannot overwrite a game, even though it's unlikely
	err = mock.CreateGame(newGame)
	assert.NotEqual(t, nil, err)
	err = mock.CreateGame(models.Game{})
	assert.NotEqual(t, nil, err)
}

func TestDeleteGame(t *testing.T) {
	mock := Mock{}
	mock.games = make(map[uuid.UUID]models.Game)
	gameID := uuid.New()
	userID := uuid.New()
	newGame := models.Game{
		ID:               gameID,
		Title:            "NewGame",
		Description:      "lame",
		UserID:           userID,
		VersionNumber:    "1.0.0",
		PublishTimestamp: "7/29/2022",
	}
	mock.games[gameID] = newGame
	err := mock.DeleteGame(gameID)
	assert.Equal(t, nil, err)
	var id uuid.UUID
	err = mock.DeleteGame(id)
	assert.NotEqual(t, nil, err)
	err = mock.DeleteGame(uuid.New())
	assert.NotEqual(t, nil, err)
}

func TestUpdateGame(t *testing.T) {
	mock := Mock{}
	mock.games = make(map[uuid.UUID]models.Game)
	gameID := uuid.New()
	userID := uuid.New()
	newGame := models.Game{
		ID:               gameID,
		Title:            "NewGame",
		Description:      "lame",
		UserID:           userID,
		VersionNumber:    "1.0.0",
		PublishTimestamp: "7/29/2022",
	}
	mock.games[gameID] = newGame
	updatedTitle := "NewerGame"
	newGame.Title = updatedTitle
	err := mock.UpdateGame(newGame)
	assert.Contains(t, mock.games, gameID)
	assert.Equal(t, mock.games[gameID].Title, updatedTitle)
	assert.Equal(t, nil, err)
	newGame.ID = uuid.New()
	err = mock.UpdateGame(newGame)
	assert.NotEqual(t, nil, err)
	err = mock.UpdateGame(models.Game{})
	assert.NotEqual(t, nil, err)
}

func TestGetAllUsers(t *testing.T){
	mock := Mock{}
	mock.Connect()
	mock.users = make([]models.User, 0)
	if(len(mock.users) > 0){
		allUsers, _ := mock.GetAllUsers()
		assert.Equal(t, len(mock.users), len(allUsers))
	}
	
}

func TestGetUserByID(t *testing.T) {
	mock := Mock{}
	mock.Connect()
	mock.users = make([]models.User, 0)
	if(len(mock.users) > 0){
		keys := reflect.ValueOf(mock.users).MapKeys()
		id := keys[rand.Intn(len(keys))].Interface().(uuid.UUID)
		user, _ := mock.GetUserByID(id)
		assert.Equal(t, user.ID, id)
	}
}

func TestCreateUser(t *testing.T){
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	userID := uuid.New()
	Role := 1
	newUser := models.User{

		ID: userID,
		Username: "testUser",
		EmailAddress "user@user.com",
		Password     "testPassword",
		Displayname  "testDisplayName",
		Role: 1
	}
	err := mock.CreateUser(newUser)
	assert.Contains(t, mock.users, userID)
	assert.Equal(t, nil, err)
	err = mock.CreateUser(newUser)
	assert.NotEqual(t, nil, err)
	err = mock.CreateUser(models.User{})
	assert.NotEqual(t, nil, err)
}

func TestDeleteUser(t *testing.T){
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	userID := uuid.New()
	Role := 1
	newUser := models.User{

		ID: userID,
		Username: "testDeleteUser",
		EmailAddress "user@user.com",
		Password     "testPassword",
		Displayname  "testDisplayName",
		Role: 2
	}
	err := mock.CreateUser(newUser)
	assert.Contains(t, mock.users, userID)
	assert.Equal(t, nil, err)
	err = mock.CreateUser(newUser)
	assert.NotEqual(t, nil, err)
	err = mock.CreateUser(models.User{})
	assert.NotEqual(t, nil, err)
}