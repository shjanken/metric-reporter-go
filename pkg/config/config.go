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

// Output object
type Output struct {
	Haicj string
	Esclt string
}

// Analysis 代表配置文了里面得 `analysis` 配置项目
type Analysis struct {
	Output *Output
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
			Output: &Output{
				Haicj: viper.GetString("analysis.output.haicj"),
				Esclt: viper.GetString("analysis.output.esclt"),
			},
		},
	}, nil
}

// Path config 的配置文件设置
type path struct {
	paths []string
}

// NewProductPath 创建生产环境的 config 文件读取路径
func NewProductPath() ConfReader {
	return &path{
		[]string{
			"/etc/metric-reporter/config",
			"$HOME/.metric-reporter/config",
			".",
		},
	}
}

// NewTestPath 创建测试环境的 config 文件读取路径
func NewTestPath() ConfReader {
	return &path{
		[]string{
			"../../configs",
		},
	}
}

// ConfReader 创建生产环境的 config 文件的读取位置
func (p *path) ReadInConfig() error {
	for _, p := range p.paths {
		viper.AddConfigPath(p)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("metric")

	return viper.ReadInConfig()
}
