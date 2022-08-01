package models

import "time"

type Game struct {
	Name         string
	Author       string
	CreationDate time.Time
	Version      string
	Tags         []string
	downloads    uint64
}
