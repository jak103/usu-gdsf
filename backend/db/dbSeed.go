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
    "TimesPlayed": 400,
    "ImagePath": "../../frontend/src/assets/minecraft.png",
    "Description": "Explore some blocks",
    "Developer": "Mojang",
    "CreationDate": "2022-08-09T01:22:30Z",
    "Version": "1.x",
    "Tags": ["Sandbox", "Adventure"],
    "downloads": 160000000
  },
  {
    "Id": "2",
    "Name": "Call of Duty Warzone",
    "Rating": 2.0,
    "TimesPlayed": 1200,
    "ImagePath": "../../frontend/src/assets/warzone.png",
    "Description": "Exciting First Person Shooter",
    "Developer": "Activision",
    "CreationDate": "2008-05-02T15:04:05Z",
    "Version": "1.x",
    "Tags": ["Spring2020", "FPS"],
    "downloads": 1000000
  },
  {
    "Id": "3",
    "Name": "Rust",
    "Rating": 3.5,
    "TimesPlayed": 400,
    "ImagePath": "../../frontend/src/assets/rust.png",
    "Description": "Survive among the rest",
    "Developer": "Facepunch Studios Ltd",
    "CreationDate": "2012-12-07T15:04:05Z",
    "Version": "1.x",
    "Tags": ["Survival", "Adventure", "Spring2018"],
    "downloads": 1000000000
  },
  {
    "Id": "4",
    "Name": "Among Us",
    "Rating": 5.0,
    "TimesPlayed": 70,
    "ImagePath": "../../frontend/src/assets/minecraft.png",
    "Description": "Work with your team to complete tasks. However, pay attention because your being hunted.",
    "Developer": "InnerSloth LLC",
    "CreationDate": "2001-05-02T15:04:05Z",
    "Version": "1.x",
    "Tags": ["Teamwork"],
    "downloads": 1060989000
  },
  {
    "Id": "5",
    "Name": "Fall guys",
    "Rating": 4.8,
    "TimesPlayed": 36,
    "ImagePath": "../../frontend/src/assets/fallguys.png",
    "Description": "Complete silly maps and puzzles with friends",
    "Developer": "Mediatonic",
    "CreationDate": "2020-08-04T15:04:05Z",
    "Version": "1.x",
    "Tags": ["Fall2020", "Adventure"],
    "downloads": 1234567891011
  }
]`
