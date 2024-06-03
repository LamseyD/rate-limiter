package redisrules

import (
	"ratelimiter/pkg/repository/redis"

	"go.uber.org/zap"
)

type Rule struct {
	RefillRate        int
	DepletionRate     int
	TemporarilyBanned bool
}

type RedisRulesRepository interface {
	GetRules(ruleType string) (Rule, error)
}

type redisRulesRepository struct {
	logger *zap.Logger
	conf   redis.Config
}

// GetRules implements RedisRulesRepository.
func (r *redisRulesRepository) GetRules(ruleType string) (Rule, error) {
	panic("unimplemented")
}

func NewRedisRulesRepository(logger *zap.Logger, conf redis.Config) RedisRulesRepository {
	return &redisRulesRepository{
		logger: logger,
		conf:   conf,
	}
}
