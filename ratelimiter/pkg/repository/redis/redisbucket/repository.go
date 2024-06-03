package redisbucket

import (
	"context"
	"ratelimiter/pkg/repository/redis"

	"go.uber.org/zap"
)

type UserBucket struct {
	ClientID   string
	RefillRate int
	Count      int
	Quota      int
}

type RedisBucketsRepository interface {
	// For TokenBucket Algorithm
	StoreBucket(ctx context.Context, bucket UserBucket) error
	GetBucket(ctx context.Context, clientID string) (UserBucket, error)
	GetAllBuckets(ctx context.Context) ([]UserBucket, error)
}

type redisBucketsRepository struct {
	logger *zap.Logger
	conf   redis.Config
}

// GetAllBuckets implements RedisBucketsRepository.
func (r *redisBucketsRepository) GetAllBuckets(ctx context.Context) ([]UserBucket, error) {
	panic("unimplemented")
}

// GetBucket implements RedisBucketsRepository.
func (r *redisBucketsRepository) GetBucket(ctx context.Context, clientID string) (UserBucket, error) {
	panic("unimplemented")
}

// StoreBucket implements RedisBucketsRepository.
func (r *redisBucketsRepository) StoreBucket(ctx context.Context, bucket UserBucket) error {
	panic("unimplemented")
}

func NewRedisBucketsRepository(logger *zap.Logger, conf redis.Config) RedisBucketsRepository {
	return &redisBucketsRepository{
		logger: logger,
		conf:   conf,
	}
}
