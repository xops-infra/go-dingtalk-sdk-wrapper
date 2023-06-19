package go_dingtalk_sdk_wrapper

import (
	"context"
	"testing"
)

func TestRobot(t *testing.T) {
	config := DingTalkConfig{
		AppKey:    "xx",
		AppSecret: "xx",
	}
	client, err := NewDingTalkClient(&config)
	if err != nil {
		t.Error(err)
	}
	client.WithRobotClient()

	req := SendMessageRequest{
		AccessToken: "6962dd0a0510e905c3c40c015a72f171c2d641209a87eccc7b4f82155cba0176",
		MessageContent: MessageContent{
			MsgType: "text",
			Text: TextBody{
				Content: "费用告警",
			},
		},
	}
	err = client.RobotSvc().SendMessage(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}

}
