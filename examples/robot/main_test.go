package go_dingtalk_sdk_wrapper

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/patsnapops/go-dingtalk-sdk-wrapper"
)

var client *DingTalkClient

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	config := DingTalkConfig{
		AppKey:    os.Getenv("CLIENT_ID"),
		AppSecret: os.Getenv("CLIENT_SECRET"),
	}
	client, err = NewDingTalkClient(&config)
	if err != nil {
		panic(err)
	}
	client.WithRobotClient()
}

func TestRobot(t *testing.T) {
	req := SendMessageRequest{
		AccessToken: os.Getenv("DingTalkRobotToken"),
		MessageContent: MessageContent{
			MsgType: "text",
			Text: TextBody{
				Content: "任务 是个@的测试",
			},
			At: AtBody{
				IsAtAll: true,
			},
		},
	}
	err := client.RobotSvc().SendMessage(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
}

func TestDownloadMsgFile(t *testing.T) {
	token := client.AccessToken
	url, err := client.RobotSvc().GetDownloadMessageFileUrl(context.Background(),
		token.Token,
		"48ECiZIRMGiDlluFVo6tixz+JMgOlMqb/esv5UWK32QMCU9KqTlg8sRxRvfbruIpGvPs7DZQ3+xIvf0Pw3dhQkP7Pds7qU54oMhQytAfa+ABQ0RUv1/8gxmIbrMmcz9fUGPANkPP00xZhsf4XegmqFKMA8mQ7VEXkpKS0KmQ86kFx/ZwiyJYxO+nLm7eskk0mDMvjU9JMAaQU7ZCYrFmrey2m2cViZtNZtel6bFIYuM=",
		"dingtc0uerrekoaiw5hn")
	if err != nil {
		t.Error(err)
	}
	t.Log(*url)
}
