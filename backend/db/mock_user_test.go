package db

import (
	"testing"

	"github.com/jak103/usu-gdsf/models"
	"github.com/stretchr/testify/assert"
)

func TestMock1(t *testing.T) {
	mock := Mock{}
	mock.Connect()
}

func TestCreateUser(t *testing.T) {
	mock := Mock{}
	mock.Connect()

	user := models.User{Username: "Rohit", Displayname: "Rohit M"}
	user.SetUUID()
	mock.CreateUser(user)

	getuser, _ := mock.GetUserByID(user.ID)

	assert.Equal(t, user.ID, getuser.ID)
}
