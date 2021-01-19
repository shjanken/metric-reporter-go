package ding

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBot(t *testing.T) {
	Convey("NewBot func should return a bot with default value", t, func() {
		bot := NewBot("fake_url", "fake_token", "fake_tmpl")

		So(bot.sendFunc, ShouldNotBeNil)
		So(bot.MsgType, ShouldEqual, Text)
		So(bot.Tmpl, ShouldEqual, "fake_tmpl")
		So(bot.Token, ShouldEqual, "fake_token")
		So(bot.URL, ShouldEqual, "fake_url")
	})
}

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

func TestCreateMsg(t *testing.T) {
	Convey("test create robot msg", t, func() {
		Convey("create msg should return a robotmsg struct", func() {
			r, err := createMsg(Text, "# hello world")
			msg, ok := r.(*robotTextMsg)

			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)
			So(msg.MsgType, ShouldEqual, "text")
			So(msg.Text.Content, ShouldEqual, "# hello world")
		})
	})
}

func TestReportMetric(t *testing.T) {
	Convey("test ReportMetric func", t, func() {
		Convey("should send request after invoke ReportMetric", func() {
			check := struct {
				result  map[string]string
				invoked bool
			}{}

			bot := NewBot("http://fake_url", "fake_token", "fake_tmpl")
			bot.sendFunc = func(url string, data []byte) error {
				check.invoked = true
				check.result = make(map[string]string)
				check.result["url"] = url
				check.result["data"] = string(data)

				return nil
			}

			bot.ReportMetric("fake_data")
			// fmt.Println(check)

			So(check.invoked, ShouldBeTrue)
			So(check.result["url"], ShouldEqual, "http://fake_url?access_token=fake_token")
		})
	})
}
