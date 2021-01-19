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

func TestCreateMsg(t *testing.T) {
	Convey("test create robot msg", t, func() {
		Convey("create msg should return a robotmsg struct", func() {
			tmpl := `# {{.Title}} {{.Content}}`
			data := struct {
				Title   string
				Content string
			}{
				Title:   "hello",
				Content: "world",
			}

			r, err := createMsg(Text, tmpl, data)
			msg, ok := r.(*robotTextMsg)

			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)
			So(msg.MsgType, ShouldEqual, "text")
			So(msg.Text.Content, ShouldContainSubstring, "# hello world")
		})
	})
}
