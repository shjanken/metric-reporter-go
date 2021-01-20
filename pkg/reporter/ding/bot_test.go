package ding

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRenderTmpl(t *testing.T) {
	Convey("test rendTmpl func", t, func() {
		temp := `### {{.Title}}
msg: {{.Content}}`

		data := struct {
			Title   string
			Content string
		}{
			Title:   "hello",
			Content: "world",
		}

		Convey("render template should success", func() {

			var b bytes.Buffer
			err := renderTmpl(&b, temp, data)

			So(err, ShouldBeNil)
			So(b.String(), ShouldEqual, "### hello\nmsg: world")
		})
	})
}

func TestReportMetric(t *testing.T) {
	Convey("test ReportMetric func", t, func() {
		check := struct {
			invoked bool
			url     string
			data    []byte
		}{}

		Convey("should send text msg after invoke ReportMetric", func() {

			markdownMsg, err := NewMarkdonwMsg("fake title", `{{.Msg}}`, struct {
				Msg string
			}{
				Msg: "fake msg",
			})

			bot := Bot{
				URL:   "http://fake_url",
				Token: "fake_token",
				Msg:   markdownMsg,
			}
			// bot := NewMarkdonwMsg("http://fake_url", "fake_token", "fake_tmpl")
			bot.sendFunc = func(url string, data []byte) error {
				check.invoked = true
				check.url = url
				check.data = data

				return nil
			}

			bot.ReportMetric()

			So(err, ShouldBeNil)
			So(check.invoked, ShouldBeTrue)
			So(check.url, ShouldEqual, "http://fake_url?access_token=fake_token")
			So(string(check.data), ShouldEqual, `{"markdown":{"text":"fake msg","title":"fake title"},"msgtype":"markdown"}`)
		})
		Convey("should send actiondCard msg after invoke ReportMetric", func() {

			markdownMsg, err := NewActionCardMsg("fake title", "fake_link", `{{.Msg}}`, struct {
				Msg string
			}{
				Msg: "fake msg",
			})

			bot := Bot{
				URL:   "http://fake_url",
				Token: "fake_token",
				Msg:   markdownMsg,
			}
			// bot := NewMarkdonwMsg("http://fake_url", "fake_token", "fake_tmpl")
			bot.sendFunc = func(url string, data []byte) error {
				check.invoked = true
				check.url = url
				check.data = data

				return nil
			}

			bot.ReportMetric()

			So(err, ShouldBeNil)
			So(check.invoked, ShouldBeTrue)
			So(check.url, ShouldEqual, "http://fake_url?access_token=fake_token")
			So(string(check.data), ShouldEqual, `{"actionCard":{"singleTitle":"查看详细","singleURL":"fake_link","text":"fake msg","title":"fake title"},"msgtype":"actionCard"}`)
		})
	})
}
