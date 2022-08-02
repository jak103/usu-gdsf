package models

import "time"

type Game struct {
	Id           int
	Name         string
	Author       string
	CreationDate time.Time
	Version      string
	Tags         []string
	downloads    uint64
}
