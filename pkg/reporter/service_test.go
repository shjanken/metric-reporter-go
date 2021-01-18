package reporter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type fake struct {
	invoke bool
}

func (f *fake) ReadMetric() (interface{}, error) {
	f.invoke = true
	return nil, nil
}

func (f *fake) ReportMetric(metric interface{}) error {
	f.invoke = true
	return nil
}

func TestGetMetric(t *testing.T) {
	Convey("Test GetMetric func", t, func() {
		Convey("should invoke the MetricProvider's func", func() {
			f := fake{false}
			_, err := GetMetric(&f)

			So(err, ShouldBeNil)
			So(f.invoke, ShouldBeTrue)
		})
	})
}

func TestReportMetric(t *testing.T) {
	Convey("Test ReportMetric func", t, func() {
		Convey("should invoke Reporter interface func", func() {
			f := &fake{false}

			err := ReportMetric(f, nil)

			So(err, ShouldBeNil)
			So(f.invoke, ShouldBeTrue)
		})
	})
}
