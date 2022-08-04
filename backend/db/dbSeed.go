package db

import (
	"encoding/json"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
)

func CreateGamesFromJson() []models.Game {
	result := []models.Game{}
	err := json.Unmarshal([]byte(JSON_SEED_DATA), &result)

	if err != nil {
		log.Warn("Unable to seed example data.")
		return nil
	}

	return result
}

const JSON_SEED_DATA = `[
  {
    "Id": "1",
    "Name": "Minecraft",
    "Rating": 4.5,
    "TimesPlayed": 40000000,
    "ImagePath": "../../frontend/src/assets/minecraft.png",
    "Description": "Explore some blocks",
    "Developer": "Mojang",
    "CreationDate": "Friday May 1, 2009",
    "Version": "1.x",
    "Tags": ["Sandbox", "Adventure"],
    "downloads": 160000000
  }
]`
