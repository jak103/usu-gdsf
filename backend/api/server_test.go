package api

import (
	"sync"
)

var GlobalTestServer Server = *NewServer(&sync.WaitGroup{})

func init() {
	go GlobalTestServer.Start()
}
