package ngxlogs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/shjanken/metric_reporter/pkg/ngxlogs/domain"
)

// Sender 接口把监控项目发的到真正的报告后端
type Sender interface {
	SendReport(m *domain.Metric) error
}

// Destination 是报告的目的地
type Destination interface {
	Sender
}

// Reporter 可以从 io.Reader 里面读一个 json 内容，并解析。
// 将解析过的监控项目发送到 Destination
type Reporter interface {
	ReadMetric() (*domain.Metric, error)
	Send(m *domain.Metric) error
}

// Service is report service
type service struct {
	r io.Reader
	d Destination
}

// New 创建一个新的对象，可以读取和发送监控项目
func New(r io.Reader, d Destination) Reporter {
	return &service{
		r, d,
	}
}

// ReadMetric 从 reader 里面读取需要报告的监控项目
func (s *service) ReadMetric() (m *domain.Metric, err error) {
	if s.r == nil {
		return nil, errors.New("can not read data, io.Reader is nil")
	}
	decode := json.NewDecoder(s.r)

	if err = decode.Decode(&m); err != nil {
		return nil, fmt.Errorf("decode json failure. %w", err)
	}

	return m, nil
}

// Send the metric to the sender backend
func (s *service) Send(m *domain.Metric) error {
	if s.d == nil {
		return errors.New("can not write data. Destination is nil")
	}
	return s.d.SendReport(m)
}
