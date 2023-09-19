package go_dingtalk_sdk_wrapper

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
)

func TestRobot(t *testing.T) {
	config := DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	}
	client, err := NewDingTalkClient(&config)
	if err != nil {
		t.Error(err)
	}
	client.WithRobotClient()
	req := SendMessageRequest{
		AccessToken: "xxxx",
		MessageContent: MessageContent{
			MsgType: "text",
			Text: TextBody{
				Content: "任务 是个@的测试",
			},
			At: AtBody{
				AtUserIDS: []string{"29070242122092575562"},
			},
		},
	}
	fmt.Println(tea.Prettify(req.MessageContent))
	err = client.RobotSvc().SendMessage(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}

}
