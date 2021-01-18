package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadConfig(t *testing.T) {
	Convey("test config pakcage", t, func() {
		Convey("test read config with test ReadInConf", func() {
			testConf := NewTestPath()

			conf, err := ReadConfig(testConf)

			if err != nil {
				t.Fatalf("read conf file failure %v", err)
			}
			So(conf.Analysis.Output.Haicj, ShouldEqual, "/home/shjanken/nginx-logs/output/haicj/")
			So(conf.Esclt.Dest, ShouldEqual, "/home/shjanken/nginx-logs/esclt")
		})
		Convey("test read product config", func() {
			productConf := NewProductPath()

			conf, err := ReadConfig(productConf)

			if err != nil {
				t.Fatalf("read conf file failure %v", err)
			}

			So(conf.Analysis.Output.Haicj, ShouldEqual, "/home/shjanken/nginx-logs/output/haicj/")
			So(conf.Esclt.Dest, ShouldEqual, "/home/shjanken/nginx-logs/esclt")
		})
	})
}
