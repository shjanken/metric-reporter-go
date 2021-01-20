package ngx

import (
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadMetric(t *testing.T) {
	Convey("test provider ReadMetric func", t, func() {
		Convey("read metric from string", func() {
			data := `{
				"general": {
					"valid_requests": 100,
					"unique_visitors": 120,
					"bandwidth": 10
				}	
			}`
			fakeProvider := MetricProvider{
				strings.NewReader(data),
			}

			res, err := fakeProvider.ReadMetric()
			m, ok := res.(Metric)
			if !ok {
				t.Fatalf("convert to metric failure")
			}

			So(err, ShouldBeNil)
			So(m.General.Request, ShouldEqual, 100)
			So(m.General.Visitors, ShouldEqual, 120)
			So(m.General.BrandWidth, ShouldEqual, 10/1024/1024)
		})

		Convey("read metric from file", func() {
			file, err := os.Open("../../../testdata/20210115.json")
			if err != nil {
				t.Fatalf("read file failure: %v", err)
			}
			provider := MetricProvider{
				file,
			}

			res, err := provider.ReadMetric()
			if err != nil {
				t.Fatalf("read file failure : %v", err)
			}
			m := res.(Metric)

			So(m.General.Request, ShouldEqual, 1048726)
			So(m.General.Visitors, ShouldEqual, 54233)
			So(m.General.BrandWidth, ShouldEqual, 19121648751/1024/1024)
		})
	})
}
