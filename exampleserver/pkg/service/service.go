package service

import (
	"context"

	"go.uber.org/zap"
)

type Service interface {
	HelloWorld(ctx context.Context) error
}

type service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) (Service, error) {
	return &service{
		logger: logger,
	}, nil
}

func (s *service) HelloWorld(ctx context.Context) error {
	s.logger.Info("Hello World")
	return nil
}
