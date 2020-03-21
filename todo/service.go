package todo

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	List(ctx context.Context) ([]*TODO, error)
	Insert(ctx context.Context, title, text string) error
}


type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     log.With(logger, "Service", "Todo"),
	}
}

func (s *service) List(ctx context.Context) ([]*TODO, error) {
	list, err := s.repository.List(ctx)
	if err != nil {
		level.Error(s.logger).Log("Service", "Todo", "List", err.Error())
		return nil, err
	}
	s.logger.Log("Service", "todo", "List", "Success")
	return list, nil
}

func (s *service) Insert(ctx context.Context, title, text string) error {
	if err := s.repository.Insert(ctx, title, text); err != nil {
		level.Error(s.logger).Log("Service", "Todo", "Insert", err.Error())
		return err
	}
	s.logger.Log("Service", "Todo", "Insert", "Success")
	return nil
}
