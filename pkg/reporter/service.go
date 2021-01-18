// Package reporter 读取监控数据文件，取得需要得数据发送报告
package reporter

// MetricProvider read the metric from src
type MetricProvider interface {
	ReadMetric() (interface{}, error)
}

// Reporter report the metrci to somewhere
type Reporter interface {
	ReportMetric(metric interface{}) error
}

// GetMetric 从 provider 里面读取数据，return structed data
func GetMetric(provider MetricProvider) (interface{}, error) {
	return provider.ReadMetric()
}

// ReportMetric 调用 reporter 的方法，发送报告到指定的目的地
func ReportMetric(reporter Reporter, metric interface{}) error {
	return reporter.ReportMetric(metric)
}
