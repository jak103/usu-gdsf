package models

type Game struct {
	Name         string
	Author       string
	CreationDate string
	Version      string
	Tags         []string
	downloads    uint64
}
