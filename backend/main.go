package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jak103/usu-gdsf/api"
	"github.com/jak103/usu-gdsf/log"
)

const version = "v0.0.0"

func main() {
	log.Info("USU Gave Dev Store Front -- %v", version)
	wg := &sync.WaitGroup{}

	server := api.NewServer(wg)
	setupCloseHandler(server)

	go server.Start()

	// Do other stuff here

	wg.Wait()
}

func setupCloseHandler(s *api.Server) {
	log.Debug("Setting up close handler")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		s.Shutdown()
	}()
}
