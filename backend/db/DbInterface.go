package db

import (
	"github.com/jak103/usu-gdsf/models"
)

// UnoDB declares the database types for the applicaiton
type StoreDatabase interface {
	// Returns all games in the database
	GetAllGameRecords() (*[]models.GameRecord, error)

	// disconnects from the database.
	disconnect()
	// connect to the database
	connect()
}
