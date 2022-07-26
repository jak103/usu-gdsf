package models

import "time"

type Game struct {
	Id           uint64
	Name         string
	Author       string
	CreationDate time.Time
	Version      string
	Tags         []string
}
