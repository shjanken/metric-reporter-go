// Package ding 包含将信息发送到 钉钉 机器人的功能
package ding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/shjanken/metric_reporter/pkg/reporter"
)

// Bot 钉钉机器的配置信息
type Bot struct {
	URL   string
	Token string
	// Tmpl     string
	Msg      *MsgData
	sendFunc func(url string, msg []byte) error
}

// MsgData 代表了要发送的钉钉机器人的消息主体
type MsgData map[string]interface{}

// NewReporter 创建一个实现了 service.Reporter 接口的对象
// url, token 是钉钉机器人的请求地址和 token
// tmpl
func NewReporter(url, token string, msg *MsgData) reporter.Reporter {
	return &Bot{
		url, token, msg, postSend,
	}
}

// NewMarkdonwMsg 构造一个用来发送给 钉钉机器人 的消息对象，最后会被编码到 json
// title, tmpl data 一起组成消息主题的内容，title 是消息的标题. 消息的 `content` 是使用 data 渲染 tmpl 模板得到的
func NewMarkdonwMsg(title, tmpl string, data interface{}) (*MsgData, error) {
	var rendered bytes.Buffer
	if err := renderTmpl(&rendered, tmpl, data); err != nil {
		return nil, err
	}

	return &MsgData{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  rendered.String(),
		},
	}, nil
}

// NewActionCardMsg 创建一个钉钉机器人的消息对象，类型是 actionCard
func NewActionCardMsg(title, link, tmpl string, data interface{}) (msg *MsgData, err error) {
	var content bytes.Buffer
	err = renderTmpl(&content, tmpl, data)

	msg = &MsgData{
		"msgtype": "actionCard",
		"actionCard": &MsgData{
			"title":       title,
			"text":        content.String(),
			"singleTitle": "查看详细报告",
			"singleURL":   link,
		},
	}

	return
}

// PostSend 调的标准库里面的 http.Post 函数发送请求。
// 为了解耦将这个函读独立出来处理
func postSend(url string, data []byte) error {
	log.Printf("Post data: %s\n", string(data))
	log.Printf("Post url: %s\n", url)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		return err
	}

	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

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

// ReportMetric 发送信息到钉钉机器人
func (b *Bot) ReportMetric() error {
	/* 	var byte bytes.Buffer
	err := renderTmpl(&byte, b.MsgData, metric)
	msg, err := createMsg(b.MsgType, byte.String()) */

	body, err := json.Marshal(*b.Msg)
	if err != nil {
		return fmt.Errorf("render template failure %w", err)
	}

	url := fmt.Sprintf("%s?access_token=%s", b.URL, b.Token)
	return b.sendFunc(url, body) // 实际调用外部服务发送消息
}
