package main

import (
	"log"

	"github.com/shjanken/metric_reporter/cmd/metric/app"
	"github.com/zloylos/grsync"
)

func main() {
	// 读取配置文件
	rConf, err := app.ReadRsyncConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err = rsync(rConf.HaicjLogSrc, rConf.HaicjLogDest); err != nil {
		log.Fatalf("rsync haicj log failure. %v", err)
	}

	if err = rsync(rConf.EscltLogSrc, rConf.EscltLogDest); err != nil {
		log.Fatalf("rsync esclt logs failure. %v", err)
	}

	// 创建用于分析的 json 文件
	if err = app.CreateJSONFile(
		[]string{
			"/home/shjanken/nginx-logs/esclt/access.log-20210111",
		},
		"/home/shjanken/nginx-logs/report.json",
	); err != nil {
		log.Fatalf("create json file failure. %v", err)
	}
}

func rsync(src, dest string) error {
	rOpt := grsync.RsyncOptions{
		Archive:  true,
		Compress: true,
		Delete:   true,
		Quiet:    true,
		Exclude:  []string{"nginx.pid", "access.log", "error.log"},
	}

	task := grsync.NewTask(
		src,
		dest,
		rOpt,
	)

	return task.Run()
}

func init() {
	log.SetPrefix("metric")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
