package ratelimiter

type Config struct {
	MaxRequestsPerMinute int `envconfig:"MAX_REQUESTS_PER_MINUTE" default:"60"`
}
