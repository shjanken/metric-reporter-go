package ngx

import (
	"encoding/json"
	"io"
)

// General metrics
type General struct {
	Request    int    `json:"valid_requests"`
	Visitors   int    `json:"unique_visitors"`
	BrandWidth int    `json:"bandwidth"`
	Date       string `json:"end_date"`
}

// Metric parsed from ngx logs
type Metric struct {
	General General `json:"general"`
}

// MetricProvider implement reporter.MetricProvider
type MetricProvider struct {
	Src io.Reader
}

// ReadMetric read json data from src, return struct object
func (provider *MetricProvider) ReadMetric() (interface{}, error) {
	decoder := json.NewDecoder(provider.Src)

	var m Metric
	if err := decoder.Decode(&m); err != nil {
		return nil, err
	}
	m.General.BrandWidth = m.General.BrandWidth / 1024 / 1024

	return m, nil
}
