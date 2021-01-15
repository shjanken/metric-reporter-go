package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shjanken/metric_reporter/cmd/metric/app"
	"github.com/shjanken/metric_reporter/pkg/config"
	"github.com/spf13/viper"
	"github.com/zloylos/grsync"
)

func main() {
	// 读取配置文件
	vConf := &viperConf{}
	conf, err := config.ReadConfig(vConf)
	if err != nil {
		log.Fatal(err)
	}

	if err = rsync(conf.Haicj.Src, conf.Haicj.Dest); err != nil {
		log.Fatalf("rsync haicj log failure. %v", err)
	}

	if err = rsync(conf.Esclt.Src, conf.Esclt.Dest); err != nil {
		log.Fatalf("rsync esclt logs failure. %v", err)
	}

	// 创建用于分析的 json 文件
	t := time.Now()
	fmt.Printf("%s/access.log-%s.gz\n", conf.Esclt.Dest, t.Format("20060102"))
	if err = app.CreateOutputFile(
		[]string{
			fmt.Sprintf("%s/access.log-%s.gz", conf.Esclt.Dest, t.Format("20060102")),
		},
		conf.Analysis.Output.Esclt,
	); err != nil {
		log.Fatalf("create json file failure. %v", err)
	}
}

type viperConf struct{}

// 初始化 viper. 读取指定目录下的配置文件
func (vConf *viperConf) ReadInConfig() (err error) {
	// ReadConfig use viper return config value
	viper.AddConfigPath("configs") // 当的 config 目录下
	viper.AddConfigPath(".")       // 当前目录下的 config.yaml

	viper.SetConfigName("metric")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
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
