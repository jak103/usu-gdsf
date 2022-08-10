package api

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAvgGameRatings(t *testing.T) {
	db, err := db.NewDatabaseFromEnv()
	if err != nil {
		assert.Fail(t, "Unable to use database")
	}

	gameID := uuid.New()
	userID := uuid.New()
	ratingID := uuid.New()

	db.CreateGame(models.Game{
		ID:               gameID,
		Title:            "Test Game",
		Description:      "This is a test game",
		UserID:           userID,
		VersionNumber:    "1.0.0",
		PublishTimestamp: time.Now().String(),
		Tags:             []string{"test", "game"},
	})

	// rating one - 3 stars
	db.CreateRating(models.GameRating{
		ID:                ratingID,
		RatingValue:       3,
		RatingDescription: "I rated this game a 3 because it was ok",
		RatingTimestamp:   time.Now().String(),
		GameId:            gameID,
		UserID:            userID,
	})

	// rating two - 5 stars
	db.CreateRating(models.GameRating{
		ID:                ratingID,
		RatingValue:       5,
		RatingDescription: "Amazing game! 5 stars!",
		RatingTimestamp:   time.Now().String(),
		GameId:            gameID,
		UserID:            userID,
	})

	gamesWithRatings, err := getAvgGameRatings()
	if err != nil {
		assert.Fail(t, "An error occurred while fetching the average of the game ratings")
	}

	avg := 0.0
	for _, game := range gamesWithRatings {
		ratings, err := db.GetRatingsByGame(game.game.ID)
		if err != nil {
			assert.Fail(t, "Unable to get ratings for %v", game.game.Title)
		}
		for _, rating := range ratings {
			avg += float64(rating.RatingValue)
		}
		avg = avg / float64(len(ratings))
	}

	assert.Equal(t, 4, avg)

}
