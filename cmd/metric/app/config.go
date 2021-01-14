package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// RsyncConfig 是一的存放了配置文件内容的结构体
type RsyncConfig struct {
	HaicjLogSrc  string
	HaicjLogDest string
	EscltLogSrc  string
	EscltLogDest string
}

// ReadConfig use viper return config value
func readConfig() (err error) {
	viper.AddConfigPath("configs")
	viper.AddConfigPath(".")

	viper.SetConfigName("metric")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}

// ReadRsyncConfig read config file and return the config object
func ReadRsyncConfig() (*RsyncConfig, error) {
	if err := readConfig(); err != nil {
		return nil, fmt.Errorf("read config file failure. %w", err)
	}

	return &RsyncConfig{
		HaicjLogSrc:  viper.GetString("rsync.haicj.src"),
		HaicjLogDest: viper.GetString("rsync.haicj.dest"),
		EscltLogSrc:  viper.GetString("rsync.esclt.src"),
		EscltLogDest: viper.GetString("rsync.esclt.dest"),
	}, nil
}
