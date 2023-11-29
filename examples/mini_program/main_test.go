package go_dingtalk_sdk_wrapper

import (
	"context"
	"os"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"

	. "github.com/xops-infra/go-dingtalk-sdk-wrapper"
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
	client.WithMiniProgramClient(cast.ToInt64(os.Getenv("AGENT_ID")))
}

// Test SendGroupNotification
func TestSendGroupNotification(t *testing.T) {
	req := SendWorkNotificationRequest{
		UseridList: tea.String("manager1,manager2"),
		ToAllUser:  tea.Bool(false),
		Msg: &MessageContent{
			MsgType: "text",
			Text: TextBody{
				Content: "this is a test for dingtalkWrap.",
			},
		},
	}
	err := client.MiniProgram.SendWorkNotification(context.Background(), &req, client.AccessToken.Token)
	if err != nil {
		t.Error(err)
	}
}

// Test SendGroupNotification
func TestSendGroupNotificationToAll(t *testing.T) {
	req := SendGroupNotificationRequest{
		ChatId: tea.String("chatxxx"),
		Msg: &MessageContent{
			MsgType: "text",
			Text: TextBody{
				Content: "this is a test for dingtalkWrap.",
			},
		},
	}
	err := client.MiniProgram.SendGroupNotification(context.Background(), &req, client.AccessToken.Token)
	if err != nil {
		t.Error(err)
	}
}
