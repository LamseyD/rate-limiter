package server

import (
	"context"
	"exampleserver/pkg/service"
	"net/http"

	"github.com/LamseyD/rate-limiter/ratelimiter/pkg/middleware/ratelimiter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vrischmann/envconfig"
	"go.uber.org/zap"
)

type Server interface {
	StartServer(e *echo.Echo)
	HelloWorld(e echo.Context) error
}

type server struct {
	service service.Service
	logger  *zap.Logger
}

func NewServer(s service.Service, logger *zap.Logger, conf Config) Server {
	return &server{
		service: s,
		logger:  logger,
	}
}

func (s *server) StartServer(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add rate limiter middleware
	var rateLimiterConfig ratelimiter.Config
	if err := envconfig.Init(&rateLimiterConfig); err != nil {
		s.logger.Fatal("failed to load rate limiter config", zap.Error(err))
	}

	rateLimiterMiddleware, err := ratelimiter.NewRateLimiter(s.logger, rateLimiterConfig)
	if err != nil {
		s.logger.Fatal("failed to create rate limiter middleware", zap.Error(err))
	}

	e.Use(rateLimiterMiddleware.RateLimit)

	g := e.Group("/api/v1")
	g.GET("/hello", s.HelloWorld)

	e.Start(":8080")
}

func (s *server) HelloWorld(e echo.Context) error {
	ctx := context.Background()
	s.service.HelloWorld(ctx)
	return e.String(http.StatusOK, "Hello World")
}
