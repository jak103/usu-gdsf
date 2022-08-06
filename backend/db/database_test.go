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
	assert.Equal(t, errors.New("DB_TYPE not set"), err)

	// env set to bad value
	os.Setenv("DB_TYPE", "bad")
	conn, err = NewDatabaseFromEnv()
	assert.Equal(t, nil, conn)
	assert.Equal(t, errors.New("unknown DB_TYPE"), err)
	os.Unsetenv("DB_TYPE")

	// env set to mock
	os.Setenv("DB_TYPE", "mock")
	conn, err = NewDatabaseFromEnv()
	assert.IsType(t, &Mock{}, conn)
	assert.Equal(t, nil, err)
	os.Unsetenv("DB_TYPE")
	connection = nil

	// env set to firestore
	os.Setenv("DB_TYPE", "firestore")
	conn, err = NewDatabaseFromEnv()
	assert.IsType(t, &Firestore{}, conn)
	assert.Equal(t, nil, err)
	os.Unsetenv("DB_TYPE")
	connection = nil

	// env set to mongo
	os.Setenv("DB_TYPE", "mongo")
	conn, err = NewDatabaseFromEnv()
	assert.IsType(t, &Mongo{}, conn)
	assert.Equal(t, nil, err)
	os.Unsetenv("DB_TYPE")
}
