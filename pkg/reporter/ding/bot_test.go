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

/* func TestCreateMsg(t *testing.T) {
	Convey("test create robot msg", t, func() {
		Convey("create msg should return a robotmsg struct", func() {
			// r, err := createMsg(Text, "# hello world")
			// msg, ok := r.(*robotTextMsg)

			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)
			So(msg.MsgType, ShouldEqual, "text")
			So(msg.Text.Content, ShouldEqual, "# hello world")
		})
	})
} */

func TestReportMetric(t *testing.T) {
	Convey("test ReportMetric func", t, func() {
		Convey("should send request after invoke ReportMetric", func() {
			check := struct {
				invoked bool
				url     string
				data    []byte
			}{}

			markdownMsg := NewMarkdonwMsg("fake title", `{{.Msg}}`, struct {
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

			So(check.invoked, ShouldBeTrue)
			So(check.url, ShouldEqual, "http://fake_url?access_token=fake_token")
			So(string(check.data), ShouldEqual, `{"markdown":{"text":"fake msg","title":"fake title"},"msgtype":"markdown"}`)
		})
	})
}
