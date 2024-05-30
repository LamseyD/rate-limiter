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
		// If we try with the algorithm, and we don't get an error
		// forward the request to the handler function
		if err := r.tokenBucket(c); err != nil {
			return next(c)
		}
		// else return error
		return echo.ErrTooManyRequests
	}
}

func (r *rateLimiter) tokenBucket(c echo.Context) error {
	// Implement rate limiting logic here
	// Do we block the source IP or do we block the client?
	// https://echo.labstack.com/docs/ip-address

	clientIP := c.RealIP()
	return nil
}
