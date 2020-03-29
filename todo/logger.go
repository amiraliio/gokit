package todo

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggerService struct {
	logger  log.Logger
	service Service
}

func NewLoggerService(logger log.Logger, service Service) Service {
	return &loggerService{
		logger:  logger,
		service: service,
	}
}

func (s *loggerService) List() (todo []*TODO, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "todo",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.service.List()
}

func (s *loggerService) Insert(title, text string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "todo",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.service.Insert(title, text)
}
