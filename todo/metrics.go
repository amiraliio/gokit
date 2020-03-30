package todo

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type metricsService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	service        Service
}

func NewMetricsService(requestCount metrics.Counter, requestLatency metrics.Histogram, service Service) Service {
	return &metricsService{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		service:        service,
	}
}

func (m *metricsService) List() ([]*TODO, error) {
	defer func(begin time.Time) {
		m.requestCount.With("method", "list").Add(1)
		m.requestLatency.With("method", "list").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return m.service.List()
}

func (m *metricsService) Insert(title, text string) error {
	defer func(begin time.Time) {
		m.requestCount.With("method", "insert").Add(1)
		m.requestLatency.With("method","insert").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return m.service.Insert(title, text)
}
