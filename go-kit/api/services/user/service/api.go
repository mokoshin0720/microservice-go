package service

import (
	"context"

	"github.com/go-kit/log"
)

type service struct {
	logger log.Logger
}

type Service interface {
	Get(ctx context.Context, id string) (string, error)
}

// NewService returns a Service with all of the expected dependencies
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) Get(ctx context.Context, id string) (string, error) {
	return "Hello, " + id, nil
}
