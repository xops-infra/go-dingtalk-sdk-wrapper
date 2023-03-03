package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"testing"
)

var (
	client *DingTalkClient
)

func init() {
	config := DingTalkConfig{
		AppKey:    "",
		AppSecret: "",
	}
	client = NewDingTalkClient(&config).WithWorkflowClient()
}

func TestDingTalkClient_SetAccessToken(t *testing.T) {
	err := client.SetAccessToken()
	if err != nil {
		fmt.Println(err)
	}
}
