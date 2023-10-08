package go_dingtalk_sdk_wrapper

// https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
import (
	"context"
	"fmt"
	"net/http"
	"time"

	robot "github.com/alibabacloud-go/dingtalk/robot_1_0"
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
	Content string `json:"content"`
}

type AtBody struct {
	IsAtAll   bool     `json:"isAtAll"`
	AtMobiles []string `json:"atMobiles"`
	AtUserIDS []string `json:"atUserIds"`
}

type LinkBody struct {
	MessageUrl string `json:"messageUrl"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	Text       string `json:"text"`
}

type MarkDownBody struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type BtnBody struct {
	ActionURL string `json:"actionURL"`
	Title     string `json:"title"`
}

type ActionCardBody struct {
	HideAvatar     string    `json:"hideAvatar"`
	BtnOrientation string    `json:"btnOrientation"`
	Single         string    `json:"singleURL"`
	SingleTitle    string    `json:"singleTitle"`
	Title          string    `json:"title"`
	Text           string    `json:"text"`
	Btns           []BtnBody `json:"btns"`
}

type FeedCard struct {
	Links []LinkBody `json:"links"`
}

type MessageContent struct {
	MsgType    string         `json:"msgtype"`
	Text       TextBody       `json:"text"`
	At         AtBody         `json:"at"`
	Link       LinkBody       `json:"link"`
	MarkDown   MarkDownBody   `json:"markdown"`
	ActionCard ActionCardBody `json:"actionCard"`
}

type SendMessageRequest struct {
	AccessToken    string         `json:"access_token"`
	Sign           string         `json:"sign"`
	MessageContent MessageContent `json:"message_content"`
}

func (c *RobotClient) SendMessage(ctx context.Context, req *SendMessageRequest) error {
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
		return err
	}
	err = c.requestBuilder.sendRequest(build, &resp)
	if err != nil {
		return err
	}
	if resp.ErrCode != 0 {
		return fmt.Errorf("%s", resp.ErrMsg)
	}
	return nil
}
