package db

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabaseFromEnv(t *testing.T) {
	// No env variable
	conn, err := NewDatabaseFromEnv()
	assert.Equal(t, nil, conn)
	assert.Equal(t, errors.New("RUN_ENV not set"), err)

	// env set to bad value
	os.Setenv("RUN_ENV", "bad")
	conn, err = NewDatabaseFromEnv()
	assert.Equal(t, nil, conn)
	assert.Equal(t, errors.New("unknown RUN_ENV"), err)
	os.Unsetenv("RUN_ENV")

	// env set to mock
	os.Setenv("RUN_ENV", "mock")
	conn, err = NewDatabaseFromEnv()
	assert.IsType(t, &Mock{}, conn)
	assert.Equal(t, nil, err)
	os.Unsetenv("RUN_ENV")
	connection = nil

	// env set to firestore
	os.Setenv("RUN_ENV", "firestore")
	conn, err = NewDatabaseFromEnv()
	assert.IsType(t, &Firestore{}, conn)
	assert.Equal(t, nil, err)
	os.Unsetenv("RUN_ENV")
	connection = nil

	// env set to mongo
	os.Setenv("RUN_ENV", "mongo")
	conn, err = NewDatabaseFromEnv()
	assert.IsType(t, &Mongo{}, conn)
	assert.Equal(t, nil, err)
	os.Unsetenv("RUN_ENV")
}
