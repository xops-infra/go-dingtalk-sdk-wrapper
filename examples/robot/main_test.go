package go_dingtalk_sdk_wrapper

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/xops-infra/go-dingtalk-sdk-wrapper"
)

var robotC *RobotClient

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	robotC = NewRobotClient()
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
				IsAtAll: false,
			},
		},
	}
	err := robotC.SendMessage(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
}
