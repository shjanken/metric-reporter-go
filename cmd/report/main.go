package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shjanken/metric_reporter/cmd/report/app"
	"github.com/shjanken/metric_reporter/pkg/reporter"
)

const usage string = `Usage: 
	report file.json msg-title`

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}

	m, err := app.ReadNgxMetric(os.Args[1])
	if err != nil {
		log.Fatalf("read ngx metric from file failure, file name: %s. err: %v", os.Args[1], err)
	}

	r, err := app.CreateNgxMetricReporter(os.Args[2], m)

	if err != nil {
		log.Fatalf("create dingding reporter failure. %v", err)
	}

	reporter.ReportMetric(r)
}
