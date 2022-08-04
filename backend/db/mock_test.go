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

func TestGetAllUsers(t *testing.T) {
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	user := models.User{
		Username:     "default",
		EmailAddress: "default@gmail.com",
		Password:     "default",
		Displayname:  "Default",
		Role:         models.Admin,
	}
	for i := 0; i < 10; i++ {
		user.ID = uuid.New()
		mock.users[user.ID] = user
	}
	allUsers, _ := mock.GetAllUsers()
	assert.Equal(t, len(mock.users), len(allUsers))
	assert.Greater(t, len(allUsers), 0)
}

func TestGetUserByID(t *testing.T) {
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	user := models.User{
		Username:     "default",
		EmailAddress: "default@gmail.com",
		Password:     "default",
		Displayname:  "Default",
		Role:         models.Admin,
	}
	for i := 0; i < 10; i++ {
		user.ID = uuid.New()
		mock.users[user.ID] = user
	}

	keys := reflect.ValueOf(mock.users).MapKeys()
	id := keys[rand.Intn(len(keys))].Interface().(uuid.UUID)
	randUser, _ := mock.GetUserByID(id)

	assert.Equal(t, randUser.ID, id)
}

func TestGetUserByRole(t *testing.T) {
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	user := models.User{
		Username:     "default",
		EmailAddress: "default@gmail.com",
		Password:     "default",
		Displayname:  "Default",
		Role:         models.Admin,
	}
	userCount := 10
	halfUsers := int(userCount / 2)
	for i := 0; i < userCount; i++ {
		user.ID = uuid.New()
		if i < halfUsers {
			user.Role = models.Publisher
		} else {
			user.Role = models.Admin
		}
		mock.users[user.ID] = user
	}
	admins, _ := mock.GetUsersByRole(0)
	publishers, _ := mock.GetUsersByRole(1)

	assert.Equal(t, halfUsers, len(admins))
	assert.Equal(t, halfUsers, len(publishers))
}

func TestCreateUser(t *testing.T) {
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	userID := uuid.New()
	newUser := models.User{
		ID:           userID,
		Username:     "default",
		EmailAddress: "default@gmail.com",
		Password:     "default",
		Displayname:  "Default",
		Role:         models.Admin,
	}
	err := mock.CreateUser(newUser)
	assert.Contains(t, mock.users, userID)
	assert.Equal(t, nil, err)
	// verify that you cannot overwrite a user, even though it's unlikely
	err = mock.CreateUser(newUser)
	assert.NotEqual(t, nil, err)
	err = mock.CreateUser(models.User{})
	assert.NotEqual(t, nil, err)
}

func TestDeleteUser(t *testing.T) {
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	userID := uuid.New()
	newUser := models.User{
		ID:           userID,
		Username:     "default",
		EmailAddress: "default@gmail.com",
		Password:     "default",
		Displayname:  "Default",
		Role:         models.Admin,
	}
	mock.users[userID] = newUser
	err := mock.DeleteUser(userID)
	assert.Equal(t, nil, err)
	var id uuid.UUID
	err = mock.DeleteUser(id)
	assert.NotEqual(t, nil, err)
	err = mock.DeleteUser(uuid.New())
	assert.NotEqual(t, nil, err)
}

func TestUpdateUser(t *testing.T) {
	mock := Mock{}
	mock.users = make(map[uuid.UUID]models.User)
	userID := uuid.New()
	newUser := models.User{
		ID:           userID,
		Username:     "default",
		EmailAddress: "default@gmail.com",
		Password:     "default",
		Displayname:  "Default",
		Role:         models.Admin,
	}
	mock.users[userID] = newUser
	updatedUsername := "username"
	newUser.Username = updatedUsername
	err := mock.UpdateUser(newUser)
	assert.Contains(t, mock.users, userID)
	assert.Equal(t, mock.users[userID].Username, updatedUsername)
	assert.Equal(t, nil, err)
	newUser.ID = uuid.New()
	err = mock.UpdateUser(newUser)
	assert.NotEqual(t, nil, err)
	err = mock.UpdateUser(models.User{})
	assert.NotEqual(t, nil, err)
}
