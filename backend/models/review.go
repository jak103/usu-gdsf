package models

type Review struct {
	Id     string
	GameId string
	UserId uint64
	Score  float64
	Text   string
}
