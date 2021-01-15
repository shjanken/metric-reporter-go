package reporter

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type MockDestionatin struct {
	invoke bool
}

func (mock *MockDestionatin) SendReport(m *Metric) error {
	mock.invoke = true
	return nil
}

func TestReadMetric(t *testing.T) {
	Convey("test ReadMetric functon", t, func() {
		Convey("read json from string", func() {
			json := `
			{
				"general": {
					"start_date": "14/Jan/2021",
					"end_date": "15/Jan/2021",
					"date_time": "2021-01-15 10:19:10 +0800",
					"total_requests": 1048734,
					"valid_requests": 1048726,
					"failed_requests": 8,
					"generation_time": 31,
					"unique_visitors": 54233,
					"unique_files": 183012,
					"excluded_hits": 0,
					"unique_referrers": 0,
					"unique_not_found": 1468,
					"unique_static_files": 4624,
					"log_size": 0,
					"bandwidth": 19121648751,
					"log_path": [
						"STDIN"
					]
				}
			}
			`
			s := New(strings.NewReader(json), nil)

			m, err := s.ReadMetric()

			So(err, ShouldBeNil)
			So(m.Requests, ShouldEqual, 1048726)
			So(m.Bandwidth, ShouldEqual, 19121648751)
			So(m.Visitors, ShouldEqual, 54233)
		})

		Convey("test send function", func() {
			mockDest := &MockDestionatin{false}
			s := New(nil, mockDest)

			err := s.Send(nil)

			So(err, ShouldBeNil)
			So(mockDest.invoke, ShouldBeTrue)
		})
	})
}
