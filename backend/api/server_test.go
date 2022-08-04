package api

import (
	"sync"
)

var GlobalTestServer Server = *NewServer(&sync.WaitGroup{})

func init() {
	println("Ran Server Test Init")
	GlobalTestServer.Start()
}
