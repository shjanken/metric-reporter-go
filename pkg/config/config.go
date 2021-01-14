package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Rsync 代表了配置文件里面 rsync 的配置项目
type Rsync struct {
	Src  string
	Dest string
}

type output struct {
	Haicj string
	Esclt string
}

// Analysis 代表配置文了里面得 `analysis` 配置项目
type Analysis struct {
	output *output
}

// Config 是一的存放了配置文件内容的结构体
type Config struct {
	Haicj    *Rsync
	Esclt    *Rsync
	Analysis *Analysis
}

// ConfReader 初始化 viper 对象。最后调用 viper.ReadInConfig 读取配置文件
type ConfReader interface {
	ReadInConfig() error
}

// ReadConfig read config file and return the config object
func ReadConfig(confReader ConfReader) (*Config, error) {
	if err := confReader.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file failure. %w", err)
	}

	return &Config{
		Haicj: &Rsync{
			Src:  viper.GetString("rsync.haicj.src"),
			Dest: viper.GetString("rsync.haicj.dest"),
		},
		Esclt: &Rsync{
			Src:  viper.GetString("rsync.esclt.src"),
			Dest: viper.GetString("rsync.esclt.dest"),
		},

		Analysis: &Analysis{
			output: &output{
				Haicj: viper.GetString("analysis.output.haicj"),
				Esclt: viper.GetString("analysis.output.esclt"),
			},
		},
	}, nil
}
