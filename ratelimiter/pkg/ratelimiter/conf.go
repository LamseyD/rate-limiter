package ratelimiter

import "ratelimiter/pkg/repository/redis"

type Config struct {
	MaxRequestsPerMinute int `envconfig:"MAX_REQUESTS_PER_MINUTE" default:"60"`

	redis.Config
}
