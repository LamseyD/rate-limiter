package requests

import (
	"exampleserver/pkg/repository/redis"

	"go.uber.org/zap"
)

type RedisRequestsRepository interface {
}

type redisRequestsRepository struct {
	logger *zap.Logger
	conf   redis.Config
}

func NewRedisRequestsRepository(logger *zap.Logger, conf redis.Config) RedisRequestsRepository {
	return &redisRequestsRepository{
		logger: logger,
		conf:   conf,
	}
}
