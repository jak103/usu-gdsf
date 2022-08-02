package models

import "time"

type Game struct {
	Id           string
	Name         string
	Author       string
	CreationDate time.Time
	Version      string
	Tags         []string
	downloads    uint64
}
