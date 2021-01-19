package main

import (
	"fmt"
	"log"
	"os"

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
		readMetricData(f)
	}

}

func readMetricData(fileName string) {
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
	if ok {
		fmt.Printf("%+v\n", m)
	}
}
