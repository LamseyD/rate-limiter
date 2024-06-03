package ratelimiter

import (
	"context"
	"errors"
	"ratelimiter/pkg/repository/redis/redisbucket"
	"ratelimiter/pkg/repository/redis/redisrules"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RateLimiter interface {
	RateLimit(next echo.HandlerFunc) echo.HandlerFunc
}

type rateLimiter struct {
	logger           *zap.Logger
	conf             Config
	rulesRepository  redisrules.RedisRulesRepository
	bucketRepository redisbucket.RedisBucketsRepository
}

func NewRateLimiter(
	logger *zap.Logger,
	conf Config,
	rulesRepository redisrules.RedisRulesRepository,
	bucketRepository redisbucket.RedisBucketsRepository,
) (RateLimiter, error) {

	return &rateLimiter{
		logger:           logger,
		conf:             conf,
		rulesRepository:  rulesRepository,
		bucketRepository: bucketRepository,
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
	// TODO Do we block the source IP or do we block the client?
	// https://echo.labstack.com/docs/ip-address
	clientIP := c.RealIP()
	ctx := context.Background()

	// TODO What is the clientID here?
	clientID := clientIP

	// TODO How do we get user rules here?
	rule, err := r.rulesRepository.GetRules("")
	if err != nil {
		return err
	}

	if rule.TemporarilyBanned {
		return errors.New("temporarily banned")
	}

	// Process rules?
	r.logger.Info("Logging rule", zap.Any("rule", rule))

	bucket, err := r.bucketRepository.GetBucket(ctx, clientID)
	if err != nil {
		return err
	}

	if bucket.Count == 0 {
		return errors.New("bucket is empty")
	}
	bucket.Count -= rule.DepletionRate
	r.bucketRepository.StoreBucket(ctx, bucket)

	return nil
}
