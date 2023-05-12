package go_dingtalk_sdk_wrapper

import (
	"context"
	"fmt"
	robot "github.com/alibabacloud-go/dingtalk/robot_1_0"
	"net/http"
	"time"
)

type RobotClient struct {
	client         *robot.Client
	requestBuilder requestBuilder
}

func NewRobotClient(client *robot.Client, builder requestBuilder) *RobotClient {
	return &RobotClient{
		client:         client,
		requestBuilder: builder,
	}
}

type TextBody struct {
	Content string `json:"content,omitempty"`
}

type AtBody struct {
	IsAtAll   bool     `json:"isAtAll,omitempty"`
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIDS []string `json:"atUserIDS,omitempty"`
}

type LinkBody struct {
	MessageUrl string `json:"messageUrl,omitempty"`
	Title      string `json:"title,omitempty"`
	PicUrl     string `json:"picUrl,omitempty"`
	Text       string `json:"text,omitempty"`
}

type MarkDownBody struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

type BtnBody struct {
	ActionURL string `json:"actionURL,omitempty"`
	Title     string `json:"title,omitempty"`
}

type ActionCardBody struct {
	HideAvatar     string    `json:"hideAvatar,omitempty"`
	BtnOrientation string    `json:"btnOrientation,omitempty"`
	Single         string    `json:"singleURL,omitempty"`
	SingleTitle    string    `json:"singleTitle,omitempty"`
	Title          string    `json:"title,omitempty"`
	Text           string    `json:"text,omitempty"`
	Btns           []BtnBody `json:"btns,omitempty"`
}

type FeedCard struct {
	Links []LinkBody `json:"links,omitempty"`
}

type MessageContent struct {
	MsgType    string         `json:"msgtype"`
	Text       TextBody       `json:"text,omitempty"`
	At         AtBody         `json:"atMobiles,omitempty"`
	Link       LinkBody       `json:"link,omitempty"`
	MarkDown   MarkDownBody   `json:"markdown,omitempty"`
	ActionCard ActionCardBody `json:"actionCard,omitempty"`
}

type SendMessageRequest struct {
	AccessToken    string         `json:"access_token"`
	Sign           string         `json:"sign"`
	MessageContent MessageContent `json:"message_content"`
}

func (c *RobotClient) SendMessage(ctx context.Context, req *SendMessageRequest) (err error) {
	var (
		resp LowApiError
		url  string
	)

	if req.Sign == "" {
		url = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", req.AccessToken)
	} else {
		timeNow := time.Now().Format("20060102")
		url = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%s&sign=%s", req.AccessToken, timeNow, req.Sign)
	}

	build, err := c.requestBuilder.build(context.Background(), http.MethodPost, url, req.MessageContent)
	if err != nil {
		return
	}
	err = c.requestBuilder.sendRequest(build, &resp)
	if err != nil {
		return
	}
	return &resp
}
