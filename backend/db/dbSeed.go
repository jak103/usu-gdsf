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
  },
  {
    "Id": "2",
    "Name": "Call of Duty Warzone",
    "Rating": 2.0,
    "TimesPlayed": 1200,
    "ImagePath": "../../frontend/src/assets/warzone.png",
    "Description": "Exciting First Person Shooter",
    "Developer": "Activision",
    "CreationDate": "Friday April 11, 2000",
    "Version": "1.x",
    "Tags": ["Spring2020", "FPS"],
    "downloads": 1000000
  },
  {
    "Id": "3",
    "Name": "Rust",
    "Rating": 3.5,
    "TimesPlayed": 40000000,
    "ImagePath": "../../frontend/src/assets/rust.png",
    "Description": "Survive among the rest",
    "Developer": "Facepunch Studios Ltd",
    "CreationDate": "Wednesday Dec 11, 2013",
    "Version": "1.x",
    "Tags": ["Survival", "Adventure", "Spring2018"],
    "downloads": 1000000000
  },
  {
    "Id": "4",
    "Name": "Among Us",
    "Rating": 5.0,
    "TimesPlayed": 7000000000,
    "ImagePath": "../../frontend/src/assets/minecraft.png",
    "Description": "Work with your team to complete tasks. However, pay attention because your being hunted.",
    "Developer": "InnerSloth LLC",
    "CreationDate": "Tuesday May 1, 2018",
    "Version": "1.x",
    "Tags": ["Teamwork"],
    "downloads": 1060989000
  },
  {
    "Id": "5",
    "Name": "Fall guys",
    "Rating": 4.8,
    "TimesPlayed": 36000000000,
    "ImagePath": "../../frontend/src/assets/fallguys.png",
    "Description": "Complete silly maps and puzzles with friends",
    "Developer": "Mediatonic",
    "CreationDate": "Tuesday Aug 4, 2020",
    "Version": "1.x",
    "Tags": ["Fall2020", "Adventure"],
    "downloads": 1234567891011
]`
