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
	URL      string
	Token    string
	Tmpl     string
	sendFunc func(url string, msg []byte) error
	MsgType  MsgType
}

type robotTextContent struct {
	Content string `json:"content"`
}

type robotTextMsg struct {
	MsgType string            `json:"msgtype"`
	Text    *robotTextContent `json:"text"`
}

// NewBot 创建一个新的钉钉机器人对象
// 默认的 msgtype 是 text (关于 msgtype , 可以查看钉钉机器人的文档)
func NewBot(url, token, tmpl string) *Bot {
	return &Bot{
		url, token, tmpl, postSend, Text,
	}
}

// PostSend 调的标准库里面的 http.Post 函数发送请求。
// 为了解耦将这个函读独立出来处理
func postSend(url string, data []byte) error {
	_, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
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

func createMsg(mType MsgType, content string) (msg interface{}, err error) {

	switch mType {
	case Text:
		msg = &robotTextMsg{
			MsgType: "text",
			Text: &robotTextContent{
				Content: content,
			},
		}
	}

	return
}

// ReportMetric 发送信息到钉钉机器人
func (b *Bot) ReportMetric(metric interface{}) error {
	var byte bytes.Buffer
	err := renderTmpl(&byte, b.Tmpl, metric)
	msg, err := createMsg(b.MsgType, byte.String())

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("render template failure %w", err)
	}

	url := fmt.Sprintf("%s?access_token=%s", b.URL, b.Token)
	return b.sendFunc(url, body) // 实际调用外部服务发送消息
}
