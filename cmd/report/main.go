package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shjanken/metric_reporter/pkg/reporter"
	"github.com/shjanken/metric_reporter/pkg/reporter/ngx"
)

const usage string = `Usage: report file.json`

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}

	for _, f := range os.Args[1:] {
		err := reportNgxMetricToDing(readNgxMetric(f))
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}
}

func reportNgxMetricToDing(metric *ngx.Metric) error {
	ddURL := os.Getenv(strings.ToUpper("dd_url"))
	ddToken := os.Getenv(strings.ToUpper("dd_token"))

	if ddURL == "" || ddToken == "" {
		return fmt.Errorf("environments DD_URL and DD_TOKEN is required")
	}

	return nil
}

func readNgxMetric(fileName string) *ngx.Metric {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("open file failure %v", err)
	}

	// data provider
	provider := ngx.MetricProvider{
		Src: file,
	}

	// read data
	r, err := reporter.GetMetric(&provider)
	if err != nil {
		log.Fatalf("read file %s failure. %v", file.Name(), err)
	}

	// convert
	var m ngx.Metric
	m, ok := r.(ngx.Metric)
	if !ok {
		log.Fatalf("convert to ngx metric failure %v", err)
	}

	return &m
}
