package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestReadConfig(t *testing.T) {
	Convey("test config pakcage", t, func() {
		cr := &testConfReader{}
		Convey("test read config file", func() {
			conf, err := ReadConfig(cr)

			So(err, ShouldBeNil)
			So(conf.Analysis.output.Haicj,
				ShouldEqual,
				"/home/shjanken/nginx-logs/")
			So(conf.Analysis.output.Esclt,
				ShouldEqual,
				"/home/shjanken/nginx-logs/")
			So(conf.Esclt.Src,
				ShouldEqual,
				"root@aliyun:/usr/local/openresty/nginx/logs")
			So(conf.Esclt.Dest,
				ShouldEqual,
				"/home/shjanken/nginx-logs/esclt")
		})
	})
}

type testConfReader struct{}

func (cr *testConfReader) ReadInConfig() error {
	// ReadConfig use viper return config value
	viper.AddConfigPath("./../../configs") // 当的 config 目录下

	viper.SetConfigName("test")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
