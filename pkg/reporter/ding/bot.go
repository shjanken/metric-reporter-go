// Package ding 包含将信息发送到 钉钉 机器人的功能
package ding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

// MsgType is dingding robot msg type.
// eg,. text, link, card ...
type MsgType int

const (
	// Text is dingding robot msg, msgtype: text
	Text MsgType = iota
	// Link is dingding robot msg, msgtype: link
	Link
)

// Bot 钉钉机器的配置信息
type Bot struct {
	URL   string
	Token string
	Temp  string
	Data  interface{}
}

type robotTextContent struct {
	Content string `json:"content"`
}

type robotTextMsg struct {
	MsgType string            `json:"msgtype"`
	Text    *robotTextContent `json:"text"`
}

func renderTmpl(writer io.Writer, t string, data interface{}) error {
	// render template
	tmpl, err := template.New("msg").Parse(t)
	if err != nil {
		return fmt.Errorf("parse template error %w", err)
	}

	if err = tmpl.Execute(writer, data); err != nil {
		return fmt.Errorf("render template failure %w", err)
	}

	return nil
}

func createMsg(mType MsgType, t string, data interface{}) (msg interface{}, err error) {
	var content bytes.Buffer
	if err = renderTmpl(&content, t, data); err != nil {
		return
	}

	switch mType {
	case Text:
		msg = &robotTextMsg{
			MsgType: "text",
			Text: &robotTextContent{
				Content: content.String(),
			},
		}
	}

	return
}

// ReportMetric 发送信息到钉钉机器人
func (b *Bot) ReportMetric(metric interface{}) error {
	robot := &robotTextMsg{
		MsgType: "text",
		Text: &robotTextContent{
			Content: "hello world",
		},
	}

	msg, err := json.Marshal(robot)
	if err != nil {
		return fmt.Errorf("Encode robot msg error. %w", err)
	}

	dest := fmt.Sprintf("%s?access_token=%s", b.URL, b.Token)
	http.Post(dest, "application/json", bytes.NewBuffer(msg))

	return nil
}
