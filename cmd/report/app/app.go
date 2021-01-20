package app

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/shjanken/metric_reporter/pkg/reporter"
	"github.com/shjanken/metric_reporter/pkg/reporter/ding"
	"github.com/shjanken/metric_reporter/pkg/reporter/ngx"
)

// ReadNgxMetric 从json文件中读取指是数据
func ReadNgxMetric(fileName string) (m *ngx.Metric, err error) {
	file, err := os.Open(fileName)

	// data provider
	provider := ngx.MetricProvider{
		Src: file,
	}

	// read data
	r, err := reporter.GetMetric(&provider)

	// convert
	metric, ok := r.(ngx.Metric)
	m = &metric
	if !ok {
		err = fmt.Errorf("convert to ngx.Metric failure. %v", err)
	}

	return
}

// CreateNgxMetricReporter 创建 reporter 对象
func CreateNgxMetricReporter(msgTitle, link string, metric *ngx.Metric) (r reporter.Reporter, err error) {
	ddURL := os.Getenv(strings.ToUpper("dd_url"))
	ddToken := os.Getenv(strings.ToUpper("dd_token"))

	if ddURL == "" || ddToken == "" {
		err = fmt.Errorf("environments DD_URL and DD_TOKEN is required")
	}

	log.Printf("%+v\n", metric)

	// 计算昨天的日期作为标题
	d := time.Now().AddDate(0, 0, -1)
	tmpl := fmt.Sprintf("%s **%s**", d.Format("2006 01月02日"), msgTitle) +
		`
- 请求数量：{{.General.Request}}
- 独立IP：  {{.General.Visitors}}
- 使用带宽：{{.General.BrandWidth}}M`

	msg, err := ding.NewActionCardMsg(msgTitle, link, tmpl, metric)
	r = ding.NewReporter(ddURL, ddToken, msg)

	// reporter.ReportMetric(r)
	return
}
