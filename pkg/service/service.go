package service

import (
	"context"
	"rate-limiter/pkg/repository/mongo"

	"go.uber.org/zap"
)

type Service interface {
	HelloWorld(ctx context.Context) error
}

type service struct {
	logger     *zap.Logger
	repository mongo.Repository
}

func NewService(logger *zap.Logger, repo mongo.Repository) (Service, error) {
	return &service{
		repository: repo,
		logger:     logger,
	}, nil
}

func (s *service) HelloWorld(ctx context.Context) error {
	s.logger.Info("Hello World")
	return nil
}
