package server

import (
	"context"
	"net/http"
	"rate-limiter/pkg/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server interface {
	StartServer(e *echo.Echo)
	HelloWorld(e echo.Context) error
}

type server struct {
	service service.Service
}

func NewServer(s service.Service) Server {
	return &server{
		service: s,
	}
}

func (s *server) StartServer(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/api/v1")
	g.GET("/hello", s.HelloWorld)

	e.Start(":8080")
}

func (s *server) HelloWorld(e echo.Context) error {
	ctx := context.Background()
	s.service.HelloWorld(ctx)
	return e.String(http.StatusOK, "Hello World")
}
