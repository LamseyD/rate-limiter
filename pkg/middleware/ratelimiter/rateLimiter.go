package ratelimiter

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RateLimiter interface {
	RateLimit(next echo.HandlerFunc) echo.HandlerFunc
}

type rateLimiter struct {
	logger *zap.Logger
	conf   Config
}

func NewRateLimiter(logger *zap.Logger, conf Config) (RateLimiter, error) {
	return &rateLimiter{
		logger: logger,
		conf:   conf,
	}, nil
}

func (r *rateLimiter) RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Implement rate limiting logic here
		return next(c)
	}
}
