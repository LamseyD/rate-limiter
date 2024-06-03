package refiller

import (
	"context"
	"math"

	"ratelimiter/pkg/repository/redis/redisbucket"
	"ratelimiter/pkg/repository/redis/redisrules"

	"go.uber.org/zap"
)

type Refiller interface {
	RefillAll(ctx context.Context) error
	RefillBucket(ctx context.Context, bucket redisbucket.UserBucket) error
}

type refiller struct {
	logger            *zap.Logger
	rulesRepository   redisrules.RedisRulesRepository
	bucketsRepository redisbucket.RedisBucketsRepository
}

// Refill implements Refiller.
func (s *refiller) RefillBucket(ctx context.Context, bucket redisbucket.UserBucket) error {
	rule, err := s.rulesRepository.GetRules("")
	if err != nil {
		return err
	}

	bucket.Count = int(math.Max(float64(bucket.Quota), float64(bucket.Count)+float64(rule.RefillRate)))

	err = s.bucketsRepository.StoreBucket(ctx, bucket)
	if err != nil {
		return err
	}

	return nil
}

// RefillAll implements Refiller.
func (s *refiller) RefillAll(ctx context.Context) error {
	buckets, err := s.bucketsRepository.GetAllBuckets(ctx)
	if err != nil {
		return err
	}

	for _, bucket := range buckets {
		err = s.RefillBucket(ctx, bucket)
		if err != nil {
			s.logger.Error(err.Error())
			continue
		}
	}
	return nil
}

func NewRefiller(
	logger *zap.Logger,
	rulesRepository redisrules.RedisRulesRepository,
	bucketRepository redisbucket.RedisBucketsRepository,
) (Refiller, error) {
	return &refiller{
		logger:            logger,
		rulesRepository:   rulesRepository,
		bucketsRepository: bucketRepository,
	}, nil
}
