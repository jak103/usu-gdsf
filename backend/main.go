package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jak103/usu-gdsf/api"
	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
)

const version = "v0.0.0"

func main() {
	log.Info("USU Gave Dev Store Front -- %v", version)
	wg := &sync.WaitGroup{}

	server := api.NewServer(wg)
	setupCloseHandler(server)

	if err := setupDatabaseConnection(); err != nil {
		log.WithError(err).Error("Error setting up database connection...")
		panic("Database not connected")
	}

	go server.Start()

	// Do other stuff here

	wg.Wait()
}

func setupDatabaseConnection() error {
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Error starting database...")
		return err
	} else {
		if cerr := db.Connect(); cerr != nil {
			log.WithError(cerr).Error("Error connecting to database...")
			return cerr
		}
	}

	return nil
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
