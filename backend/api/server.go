package api

import (
	"context"
	"sync"

	"github.com/jak103/usu-gdsf/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	wg   *sync.WaitGroup
}

func NewServer(wg *sync.WaitGroup) *Server {
	wg.Add(1)
	return &Server{
		wg: wg,
	}
}

func (s *Server) Start() {
	log.Info("Starting API server")
	s.echo = echo.New()

	s.setupMiddleware()
	s.setupRoutes()

	s.echo.Start(":8080")
}

func (s *Server) Shutdown() {
	log.Info("Shutting down API server")
	s.echo.Shutdown(context.Background())
	log.Info("Done")
	s.wg.Done()
}

func (s *Server) setupMiddleware() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Gzip())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())

	s.echo.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "/client/dist/",
		HTML5: true,
	}))
}

func (s *Server) setupRoutes() {
	s.echo.GET("/hello", hello)

}
