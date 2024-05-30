package rules

import (
	"rate-limiter/pkg/repository/redis"

	"go.uber.org/zap"
)

type RedisRulesRepository interface {
}

type redisRulesRepository struct {
	logger *zap.Logger
	conf   redis.Config
}

func NewRedisRulesRepository(logger *zap.Logger, conf redis.Config) RedisRulesRepository {
	return &redisRulesRepository{
		logger: logger,
		conf:   conf,
	}
}
