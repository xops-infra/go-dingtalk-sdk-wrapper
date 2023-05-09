package go_dingtalk_sdk_wrapper

import (
	"github.com/robfig/cron/v3"
	"testing"
)

func TestDingTalkClient_SetAccessToken(t *testing.T) {
	//fmt.Println(client)
	config := DingTalkConfig{
		AppKey:    "xx",
		AppSecret: "xx",
	}
	client, err := NewDingTalkClient(&config)
	if err != nil {
		t.Error(client)
	}
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		t.Log("获取")
		err := client.CronSetAccessToken()
		if err != nil {
			return
		}
	})
	c.Start()
	t.Log(client.AccessToken)
}
